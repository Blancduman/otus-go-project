package cmd

import (
	"context"
	"github.com/Blancduman/banners-rotation/internal/build"
	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func migrateCmd(ctx context.Context, conf config.Config) *cobra.Command {
	command := &cobra.Command{
		Use:       "migrate",
		Short:     "run db migrations",
		ValidArgs: []string{"catalog", "promo"},
		RunE: func(cmd *cobra.Command, args []string) error {
			//nolint:wrapcheck
			return cmd.Usage()
		},
	}

	command.AddCommand(up(ctx, conf))
	command.AddCommand(down(ctx, conf))

	return command
}

func up(ctx context.Context, conf config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "up migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			b := build.New(conf)

			m, err := b.CatalogMigration(ctx)
			if err != nil {
				return errors.Wrap(err, "get migrate")
			}

			err = m.Up()
			if err != nil {
				return errors.Wrap(err, "up migrations")
			}

			return nil
		},
	}
}

func down(ctx context.Context, conf config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "rollback all migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			b := build.New(conf)

			m, err := b.CatalogMigration(ctx)
			if err != nil {
				return errors.Wrap(err, "get migrate")
			}

			err = m.Down()
			if err != nil {
				return errors.Wrap(err, "down migrations")
			}

			return nil
		},
	}
}
