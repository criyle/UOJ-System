package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/criyle/UOJ-System/judger2/pb"
	exec_pb "github.com/criyle/go-judge/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // JSON grpc
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	envAddr        = "ADDR"
	envToken       = "TOKEN"
	envExecAddr    = "EXEC_ADDR"
	envExecToken   = "EXEC_TOKEN"
	envMetricsAddr = "METRICS_ADDR"
	envConfFile    = "CONF_FILE"
	envRelease     = "RELEASE"
)

var (
	addr         = ":5081"
	metricsAddr  = ":5082"
	execAddr     = ":5051"
	confFileName = "compile.yaml"
	release      = false

	token, execToken string

	logger *zap.Logger
)

func init() {
	if os.Getpid() == 1 || os.Getenv(envRelease) != "" {
		release = true
	}
	var err error
	if release {
		logger, err = zap.NewProduction()
	} else {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	}
	if err != nil {
		log.Fatalln("init logger", err)
	}

	if s := os.Getenv(envAddr); s != "" {
		addr = s
	}
	if s := os.Getenv(envExecAddr); s != "" {
		execAddr = s
	}
	if s := os.Getenv(envConfFile); s != "" {
		confFileName = s
	}

	token = os.Getenv(envToken)
	execToken = os.Getenv(envExecToken)
}

func main() {
	printLog := logger.Sugar().Info
	conf, err := readConfig(confFileName)
	if err != nil {
		logger.Fatal(err.Error())
	}
	printLog("compile config", conf)
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if execToken != "" {
		opts = append(opts,
			grpc.WithPerRPCCredentials(newTokenAuth(execToken)),
			grpc.WithChainUnaryInterceptor(
				grpc_zap.UnaryClientInterceptor(logger),
			),
			grpc.WithChainStreamInterceptor(
				grpc_zap.StreamClientInterceptor(logger),
			),
		)
	}
	conn, err := grpc.Dial(execAddr, opts...)
	if err != nil {
		log.Fatalln("client", err)
	}
	client := exec_pb.NewExecutorClient(conn)

	streamMiddleware := []grpc.StreamServerInterceptor{
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(logger),
		grpc_recovery.StreamServerInterceptor(),
	}
	unaryMiddleware := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	}
	if token != "" {
		authFunc := grpcTokenAuth(token)
		streamMiddleware = append(streamMiddleware, grpc_auth.StreamServerInterceptor(authFunc))
		unaryMiddleware = append(unaryMiddleware, grpc_auth.UnaryServerInterceptor(authFunc))
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamMiddleware...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryMiddleware...)),
	)
	pb.RegisterCompileServer(grpcServer, newServer(client, conf))
	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		printLog("Starting grpc server at", addr)
		printLog("GRPC serve", grpcServer.Serve(lis))
	}()

	// metrics
	http.Handle("/metrics", promhttp.Handler())
	srv := http.Server{
		Addr:    metricsAddr,
		Handler: http.DefaultServeMux,
	}
	go func() {
		printLog(srv.ListenAndServe())
	}()

	// Graceful shutdown...
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	printLog("Shutting Down...")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var eg errgroup.Group
	eg.Go(func() error {
		grpcServer.GracefulStop()
		printLog("GRPC server shutdown")
		return nil
	})
	eg.Go(func() error {
		printLog("Http server shutdown")
		return srv.Shutdown(ctx)
	})
	go func() {
		printLog("Shutdown Finished", eg.Wait())
		cancel()
	}()
	<-ctx.Done()
}

func grpcTokenAuth(token string) func(context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		reqToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		if reqToken != token {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}
		return ctx, nil
	}
}

type tokenAuth struct {
	token string
}

func newTokenAuth(token string) credentials.PerRPCCredentials {
	return &tokenAuth{token: token}
}

// Return value is mapped to request headers.
func (t *tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (*tokenAuth) RequireTransportSecurity() bool {
	return false
}
