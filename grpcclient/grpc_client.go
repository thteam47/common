package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type GrpcClientConn struct {
	Config  GrpcClientConnConfig
	Conn    *grpc.ClientConn
	Address string
}

type GrpcClientConnConfig struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	AccessToken string        `mapstructure:"access_token"`
}

var GrpcClientConnTimeout = 20 * time.Second

func NewGrpcClientConnWithConfig(config GrpcClientConnConfig) (*GrpcClientConn, error) {
	if config.Timeout == 0 {
		config.Timeout = GrpcClientConnTimeout
	}

	return NewGrpcClientConn(config), nil
}

func NewGrpcClientConn(config GrpcClientConnConfig) *GrpcClientConn {
	inst := &GrpcClientConn{
		Config: config,
	}

	inst.Address = inst.Config.Address

	return inst
}

func (inst *GrpcClientConn) Context() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), inst.Config.Timeout)
	defer cancel()
	return ctx
}

func (inst *GrpcClientConn) Stop() {
	if inst.Conn != nil {
		inst.Conn.Close()
		inst.Conn = nil
	}
}
