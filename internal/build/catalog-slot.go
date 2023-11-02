package build

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/catalog/slot"
)

const collectionSlot = "slot"

func (b *Builder) slotService(ctx context.Context) (*slot.Service, error) {
	repo, err := b.slotRepo(ctx)
	if err != nil {
		return nil, err
	}

	return slot.NewService(repo), nil
}

func (b *Builder) slotRepo(ctx context.Context) (*slot.RepoMongoDB, error) {
	coll, err := b.mongoCatalogCollection(ctx, collectionSlot)
	if err != nil {
		return nil, err
	}

	return slot.NewMongoDBRepo(coll), nil
}
