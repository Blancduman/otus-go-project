package build

import (
	"github.com/Blancduman/banners-rotation/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Builder struct {
	config config.Config

	mongoCatalog struct {
		client      *mongo.Client
		database    *mongo.Database
		collections map[string]*mongo.Collection
	}
}

func New(conf config.Config) *Builder {
	return &Builder{config: conf} //nolint:exhaustruct
}
