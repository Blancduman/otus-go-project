package cmd

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, conf config.Config) error {
	root := &cobra.Command{ //nolint:exhaustruct
		RunE: func(cmd *cobra.Command, args []string) error {
			//nolint:wrapcheck
			return cmd.Usage()
		},
	}

	root.AddCommand(
		grpcCmd(ctx, conf),
	)

	return errors.Wrap(root.ExecuteContext(ctx), "run application")
}
