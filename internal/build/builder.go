package build

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Blancduman/banners-rotation/internal/config"
)

type Builder struct {
	config config.Config

	mongoCatalog struct {
		client      *mongo.Client
		database    *mongo.Database
		collections map[string]*mongo.Collection
	}
}

func New(ctx context.Context, conf config.Config) *Builder {
	return &Builder{config: conf} //nolint:exhaustruct
}
