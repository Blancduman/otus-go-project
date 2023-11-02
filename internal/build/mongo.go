package build

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *Builder) mongoCatalogCollection(ctx context.Context, name string) (*mongo.Collection, error) {
	if b.mongoCatalog.collections == nil {
		b.mongoCatalog.collections = make(map[string]*mongo.Collection)
	}

	if _, ok := b.mongoCatalog.collections[name]; !ok {
		db, err := b.mongoCatalogDatabase(ctx)
		if err != nil {
			return nil, err
		}

		c := db.Collection(name)

		b.mongoCatalog.collections[name] = c
	}

	return b.mongoCatalog.collections[name], nil
}

func (b *Builder) mongoCatalogDatabase(ctx context.Context) (*mongo.Database, error) {
	if b.mongoCatalog.database == nil {
		mongoTransaction, err := b.mongoCatalogClient(ctx)
		if err != nil {
			return nil, err
		}

		b.mongoCatalog.database = mongoTransaction.Database(b.config.Mongo.DB, options.Database())
	}

	return b.mongoCatalog.database, nil
}

func (b *Builder) mongoCatalogClient(ctx context.Context) (*mongo.Client, error) {
	if b.mongoCatalog.client == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(b.config.Mongo.DSN))
		if err != nil {
			return nil, errors.Wrap(err, "connect mongodb client")
		}

		b.mongoCatalog.client = client
	}

	return b.mongoCatalog.client, nil
}
