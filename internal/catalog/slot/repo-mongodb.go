package slot

import (
	"context"

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

func (r RepoMongoDB) Get(ctx context.Context, id int64) (Slot, error) {
	var doc slotDoc

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return Slot{}, ErrNotFound
	}

	return doc.toModel(), errors.Wrap(err, "find slot by id")
}

func (r RepoMongoDB) Create(ctx context.Context, slot Slot) (int64, error) {
	var doc slotDoc
	err := r.collection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return 0, errors.Wrap(err, "get last document")
	}

	lastID := doc.ID
	slot.ID = lastID + 1

	_, err = r.collection.InsertOne(ctx, slotDocFromModel(slot))
	if err != nil {
		return 0, errors.Wrap(err, "insert document")
	}

	return slot.ID, nil
}

func (r RepoMongoDB) Update(ctx context.Context, slot Slot) (int64, error) {
	res, err := r.update(ctx, slot)

	return res, errors.Wrap(err, "update slot")
}

func (r RepoMongoDB) Delete(ctx context.Context, id int64) (int64, error) {
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, errors.Wrap(err, "delete slot by id")
	}

	return res.DeletedCount, nil
}

func (r RepoMongoDB) update(ctx context.Context, slot Slot) (int64, error) {
	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": slot.ID},
		bson.M{"$set": slotDocFromModel(slot)},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return 0, errors.Wrap(err, "upsert slot")
	}

	return res.UpsertedCount + res.ModifiedCount, nil
}
