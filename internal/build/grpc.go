package build

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
)

func (b *Builder) GRPCServer(ctx context.Context) (*grpc.Server, error) {
	logger := zerolog.Ctx(ctx)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
				l := logger.With().Fields(fields).Logger()

				switch lvl {
				case logging.LevelDebug:
					l.Debug().Msg(msg)
				case logging.LevelInfo:
					l.Info().Msg(msg)
				case logging.LevelWarn:
					l.Warn().Msg(msg)
				case logging.LevelError:
					l.Error().Msg(msg)
				default:
					panic(fmt.Sprintf("unknown level %v", lvl))
				}
			}), logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)),
		),
	)

	return grpcServer, nil
}

func (b *Builder) Listener() (net.Listener, error) {
	listener, err := net.Listen(b.config.GRPC.NetworkType, b.config.GRPCAddr())
	if err != nil {
		return nil, errors.Wrap(err, "start network listener")
	}

	return listener, nil
}
