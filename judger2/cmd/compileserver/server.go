package main

import (
	"context"
	"time"

	"github.com/criyle/UOJ-System/judger2/pb"
	exec_pb "github.com/criyle/go-judge/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const msg = "msg"

type server struct {
	pb.UnimplementedCompileServer
	client exec_pb.ExecutorClient
	config *config
}

func newServer(client exec_pb.ExecutorClient, conf *config) pb.CompileServer {
	return &server{client: client, config: conf}
}

func (s *server) Compile(ctx context.Context, req *pb.CompileRequest) (*pb.CompileResult, error) {
	lConf, ok := s.config.Language[req.GetLanguage()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Language (%s) does not found", req.GetLanguage())
	}
	// Do compile
	copyIn := make(map[string]*exec_pb.Request_File)
	copyIn[lConf.SourceFileName] = &exec_pb.Request_File{
		File: &exec_pb.Request_File_Memory{
			Memory: &exec_pb.Request_MemoryFile{
				Content: req.GetSource(),
			},
		},
	}

	execReq := &exec_pb.Request{
		Cmd: []*exec_pb.Request_CmdType{{
			Args: lConf.CompileArgs,
			Env:  lConf.CompileEnv,
			Files: []*exec_pb.Request_File{
				{File: &exec_pb.Request_File_Memory{
					Memory: &exec_pb.Request_MemoryFile{Content: []byte("")},
				}},
				{File: &exec_pb.Request_File_Pipe{
					Pipe: &exec_pb.Request_PipeCollector{Name: msg, Max: s.config.MaxCompilerMessageSize},
				}},
				{File: &exec_pb.Request_File_Pipe{
					Pipe: &exec_pb.Request_PipeCollector{Name: msg, Max: s.config.MaxCompilerMessageSize},
				}},
			},
			Tty:          true,
			CPULimit:     uint64(s.config.CompileTimeLimit),
			RealCPULimit: uint64(s.config.CompileTimeLimit + time.Second), // +1s
			MemoryLimit:  s.config.CompileMemoryLimit,
			ProcLimit:    s.config.CompileProcLimit,
			CopyIn:       copyIn,
			CopyOut:      lConf.ExecuteFileName,
			CopyOutMax:   s.config.MaxExecuteFileSize,
		}},
	}
	// do exec
	execResp, err := s.client.Exec(ctx, execReq)
	if err != nil {
		return nil, err
	}
	if err := execResp.GetError(); err != "" {
		return nil, status.Errorf(codes.Internal, err)
	}
	ret := execResp.GetResults()[0]
	if ret.Error != "" {
		return nil, status.Errorf(codes.Internal, ret.Error)
	}
	if ret.Status != exec_pb.Response_Result_Accepted {
		return &pb.CompileResult{CompileMessage: ret.Files[msg]}, nil
	}

	compilerMsg := ret.Files[msg]
	delete(ret.Files, msg)

	return &pb.CompileResult{
		Exec:           ret.Files,
		Args:           lConf.ExecuteArgs,
		Env:            lConf.ExecuteEnv,
		ProcLimit:      lConf.ExecuteProcLimit,
		CompileMessage: compilerMsg,
	}, nil
}
