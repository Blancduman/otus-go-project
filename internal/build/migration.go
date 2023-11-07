package build

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Blancduman/banners-rotation/internal/migration"
)

func (b *Builder) CatalogMigration(ctx context.Context) (*migrate.Migrate, error) {
	client, err := b.mongoCatalogClient(ctx)
	if err != nil {
		return nil, err
	}

	return b.mongoMigration(client, b.config.Mongo.DB, migration.CatalogPath, "schema_migrations")
}

func (b *Builder) mongoMigration(client *mongo.Client, database string, path string, collection string) (*migrate.Migrate, error) {
	fs, err := iofs.New(migration.FS, path)
	if err != nil {
		return nil, errors.Wrap(err, "build migrations iofs")
	}

	//nolint:exhaustruct
	db, err := mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName:         database,
		TransactionMode:      false,
		MigrationsCollection: collection,
	})
	if err != nil {
		return nil, errors.Wrap(err, "build mongo driver")
	}

	m, err := migrate.NewWithInstance("iofs", fs, database, db)
	if err != nil {
		return nil, errors.Wrap(err, "build mongo migrate instance")
	}

	return m, nil
}
