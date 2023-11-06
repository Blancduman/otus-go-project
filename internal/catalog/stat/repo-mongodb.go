package stat

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoMongoDB struct {
	collection *mongo.Collection
}

func NewMongoDBRepo(collection *mongo.Collection) *RepoMongoDB {
	return &RepoMongoDB{
		collection: collection,
	}
}

func (r RepoMongoDB) GetStat(ctx context.Context, slotID int64) (SlotStat, error) {
	var doc slotStatDoc

	err := r.collection.FindOne(ctx, bson.M{"_id": slotID}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return SlotStat{}, ErrNotFound
	}

	return doc.toModel(), errors.Wrap(err, "find stat by slot id")
}

func (r RepoMongoDB) IncrementClickedCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": slotID},
		bson.M{"$inc": bson.M{fmt.Sprintf("bannerStat.%d.%d.clicked", bannerID, socialDemGroupID): 1}},
		options.Update().SetUpsert(true),
	)

	return errors.Wrap(err, "increment clicked count")
}

func (r RepoMongoDB) IncrementShownCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": slotID},
		bson.M{"$inc": bson.M{fmt.Sprintf("bannerStat.%d.%d.shown", bannerID, socialDemGroupID): 1}},
		options.Update().SetUpsert(true),
	)

	return errors.Wrap(err, "increment shown count")
}

func (r RepoMongoDB) AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error {
	m := bson.M{}

	for _, v := range socialDemGroupIDs {
		m[fmt.Sprintf("bannerStat.%d.%d.clicked", bannerID, v)] = 0
		m[fmt.Sprintf("bannerStat.%d.%d.shown", bannerID, v)] = 0
	}

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": slotID},
		bson.M{"$set": m},
		options.Update().SetUpsert(true),
	)

	return errors.Wrap(err, "add banner to slot in stat")
}

func (r RepoMongoDB) RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": slotID},
		bson.M{"$unset": bson.M{fmt.Sprintf("bannerStat.%d", bannerID): ""}},
		options.Update(),
	)

	return errors.Wrap(err, "remove banner to slot in stat")
}
