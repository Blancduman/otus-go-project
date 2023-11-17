package cmd

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/build"
	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	grpclib "google.golang.org/grpc"
)

func grpcCmd(ctx context.Context, conf config.Config) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "grpc",
		Short: "start grpc server listening",
		RunE: func(cmd *cobra.Command, args []string) error {
			builder := build.New(conf)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			go func() {
				cancel()
			}()

			listener, err := builder.Listener()
			if err != nil {
				return errors.Wrap(err, "start network listener")
			}

			server, err := builder.ItemGRPCServer(ctx)
			if err != nil {
				return errors.Wrap(err, "build grpc server")
			}

			if err = server.Serve(listener); !errors.Is(err, grpclib.ErrServerStopped) {
				return errors.Wrap(err, "run grpc server")
			}

			<-ctx.Done()

			return nil
		},
	}
}
