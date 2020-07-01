package main

import (
	"io/ioutil"
	"time"

	"github.com/flynn/go-shlex"
	"gopkg.in/yaml.v2"
)

type config struct {
	CompileEnv             []string              `yaml:"compileEnv"`
	ExecuteEnv             []string              `yaml:"executeEnv"`
	CompileTimeLimit       time.Duration         `yaml:"compileTimeLimit"`
	CompileMemoryLimit     uint64                `yaml:"compileMemoryLimit"` // in byte
	CompileProcLimit       uint64                `yaml:"compileProcLimit"`
	ExecuteProcLimit       uint64                `yaml:"executeProcLimit"`
	MaxCompilerMessageSize int64                 `yaml:"maxCompilerMessageSize"` // in byte
	MaxExecuteFileSize     uint64                `yaml:"maxExecuteFileSize"`     // in byte
	Language               map[string]langConfig `yaml:"language"`
}

type langConfig struct {
	CompileCmd       string   `yaml:"compileCmd"`
	CompileArgs      []string `yaml:"compileArgs"`
	CompileEnv       []string `yaml:"compileEnv"`
	SourceFileName   string   `yaml:"sourceFileName"`
	ExecuteFileName  []string `yaml:"executeFileName"`
	ExecuteCmd       string   `yaml:"executeCmd"`
	ExecuteArgs      []string `yaml:"executeArgs"`
	ExecuteEnv       []string `yaml:"executeEnv"`
	ExecuteProcLimit uint64   `yaml:"executeProcLimit"`
}

func readConfig(fileName string) (*config, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	conf := new(config)
	if err := yaml.Unmarshal(b, conf); err != nil {
		return nil, err
	}
	for k, v := range conf.Language {
		if len(v.CompileArgs) == 0 {
			s, err := shlex.Split(v.CompileCmd)
			if err != nil {
				return nil, err
			}
			v.CompileArgs = s
		}
		if len(v.ExecuteArgs) == 0 {
			s, err := shlex.Split(v.ExecuteCmd)
			if err != nil {
				return nil, err
			}
			v.ExecuteArgs = s
		}
		if len(v.CompileEnv) == 0 {
			v.CompileEnv = conf.CompileEnv
		} else {
			var env []string
			env = append(env, conf.CompileEnv...)
			env = append(env, v.CompileEnv...)
			v.CompileEnv = env
		}
		if len(v.ExecuteEnv) == 0 {
			v.ExecuteEnv = conf.ExecuteEnv
		} else {
			var env []string
			env = append(env, conf.ExecuteEnv...)
			env = append(env, v.ExecuteEnv...)
			v.ExecuteEnv = env
		}
		if v.ExecuteProcLimit == 0 {
			v.ExecuteProcLimit = conf.ExecuteProcLimit
		}
		conf.Language[k] = v
	}
	return conf, nil
}
