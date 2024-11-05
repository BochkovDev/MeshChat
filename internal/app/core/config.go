package core

import (
	"time"
)

type EnvType string

const (
	Local EnvType = "local"
	Dev   EnvType = "dev"
	Prod  EnvType = "prod"
)

type Config struct {
	Debug   bool       `yaml:"debug" env-required:"true"`
	Env     EnvType    `yaml:"env" env-required:"true"`
	RootDir string     `yaml:"root_dir" env-required:"true"`
	GRPC    GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
