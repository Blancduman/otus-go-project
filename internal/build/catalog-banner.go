package build

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/catalog/banner"
)

const collectionBanner = "banner"

func (b *Builder) bannerService(ctx context.Context) (*banner.Service, error) {
	repo, err := b.bannerRepo(ctx)
	if err != nil {
		return nil, err
	}

	return banner.NewService(repo), nil
}

func (b *Builder) bannerRepo(ctx context.Context) (*banner.RepoMongoDB, error) {
	coll, err := b.mongoCatalogCollection(ctx, collectionBanner)
	if err != nil {
		return nil, err
	}

	return banner.NewMongoDBRepo(coll), nil
}
