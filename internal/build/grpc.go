package build

import (
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func (b *Builder) GRPCServer() (*grpc.Server, error) {
	grpcServer := grpc.NewServer()

	return grpcServer, nil
}

func (b *Builder) Listener() (net.Listener, error) {
	listener, err := net.Listen(b.config.GRPC.NetworkType, b.config.GRPCAddr())
	if err != nil {
		return nil, errors.Wrap(err, "start network listener")
	}

	return listener, nil
}
