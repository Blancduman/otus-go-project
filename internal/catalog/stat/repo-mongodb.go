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

//func (r RepoMongoDB) AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error {
//	m := bson.M{}
//
//	for _, v := range socialDemGroupIDs {
//		m[fmt.Sprintf("%d.%d.clicked", bannerID, v)] = 0
//		m[fmt.Sprintf("%d.%d.shown", bannerID, v)] = 0
//	}
//
//	_, err := r.collection.UpdateOne(
//		ctx,
//		bson.M{"_id": slotID},
//		bson.M{"$set": bson.M{"bannerStat": m}},
//		options.Update().SetUpsert(true),
//	)
//
//	return errors.Wrap(err, "add banner to slot in stat")
//}
//
//func (r RepoMongoDB) RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error {
//	_, err := r.collection.UpdateOne(
//		ctx,
//		bson.M{"_id": slotID},
//		bson.M{"$unset": fmt.Sprintf("bannerStat.%d.%d", slotID, bannerID)},
//		options.Update().SetUpsert(true),
//	)
//
//	return errors.Wrap(err, "remove banner to slot in stat")
//}
