package build

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/catalog/stat"
)

const collectionStat = "stat"

func (b *Builder) statService(ctx context.Context) (*stat.Service, error) {
	repo, err := b.statRepo(ctx)
	if err != nil {
		return nil, err
	}

	producer, err := b.reporterProducer()
	if err != nil {
		return nil, err
	}

	return stat.NewService(repo, producer), nil
}

func (b *Builder) statRepo(ctx context.Context) (*stat.RepoMongoDB, error) {
	coll, err := b.mongoCatalogCollection(ctx, collectionStat)
	if err != nil {
		return nil, err
	}

	return stat.NewMongoDBRepo(coll), nil
}
