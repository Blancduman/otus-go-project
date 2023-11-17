package config

import "fmt"

type GRPC struct {
	NetworkType string `envconfig:"GRPC_NETWORK_TYPE" default:"tcp"`
	Port        int32  `envconfig:"GRPC_PORT" default:"8081"`
}

func (c *Config) GRPCAddr() string {
	return fmt.Sprintf(":%d", c.GRPC.Port)
}
