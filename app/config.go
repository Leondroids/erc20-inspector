package app

import (
	"os"
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"github.com/Leondroids/gox"
)

const (
	DefaultPort          = "8080"
	ENV_PORT             = "RPC_API_PORT"
	ENV_RPC_ENDPOINT     = "RPC_ENDPOINT"
	DefaultRPCEndpoint   = "http://localhost:8545" //rpc.InfuraEndpoint
)

type Config struct {
	Port        string
	RPCEndpoint string
}

type Context struct {
	Config Config
	Client *rpc.Client
}

func InitApp() (*Context, error) {
	// config
	config := Config{
		Port:        fmt.Sprintf(":%v", gox.EnvReadStringOr(ENV_PORT, DefaultPort)),
		RPCEndpoint: gox.EnvReadStringOr(ENV_RPC_ENDPOINT, DefaultRPCEndpoint),
	}

	return &Context{
		Config: config,
		Client: rpc.NewRPCClient(config.RPCEndpoint),
	}, nil
}

