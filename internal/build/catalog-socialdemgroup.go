package build

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/catalog/socialdemgroup"
)

const collectionSocialDemGroup = "socialDemGroup"

func (b *Builder) socialDemGroupService(ctx context.Context) (*socialdemgroup.Service, error) {
	repo, err := b.socialDemGroupRepo(ctx)
	if err != nil {
		return nil, err
	}

	return socialdemgroup.NewService(repo), nil
}

func (b *Builder) socialDemGroupRepo(ctx context.Context) (*socialdemgroup.RepoMongoDB, error) {
	coll, err := b.mongoCatalogCollection(ctx, collectionSocialDemGroup)
	if err != nil {
		return nil, err
	}

	return socialdemgroup.NewMongoDBRepo(coll), nil
}
