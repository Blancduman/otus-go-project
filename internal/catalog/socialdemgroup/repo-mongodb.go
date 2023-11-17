package socialdemgroup

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

func (r RepoMongoDB) Get(ctx context.Context, id int64) (SocialDemGroup, error) {
	var doc socialDemGroupDoc

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return SocialDemGroup{}, ErrNotFound
	}

	return doc.toModel(), errors.Wrap(err, "find social dem group by id")
}

func (r RepoMongoDB) GetAll(ctx context.Context) ([]SocialDemGroup, error) {
	var docs []socialDemGroupDoc

	err := r.collection.FindOne(ctx, bson.M{}).Decode(&docs)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNotFound
	}

	res := make([]SocialDemGroup, len(docs))

	for i, d := range docs {
		res[i] = d.toModel()
	}

	return res, errors.Wrap(err, "find social dem group by id")
}

func (r RepoMongoDB) Create(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error) {
	var doc socialDemGroupDoc
	err := r.collection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return 0, errors.Wrap(err, "get last document")
	}

	lastID := doc.ID
	socialDemGroup.ID = lastID + 1

	_, err = r.collection.InsertOne(ctx, socialDemGroupDocFromModel(socialDemGroup))
	if err != nil {
		return 0, errors.Wrap(err, "insert document")
	}

	return socialDemGroup.ID, nil
}

func (r RepoMongoDB) Update(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error) {
	res, err := r.update(ctx, socialDemGroup)

	return res, errors.Wrap(err, "update social dem group")
}

func (r RepoMongoDB) Delete(ctx context.Context, id int64) (int64, error) {
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, errors.Wrap(err, "delete social dem group by id")
	}

	return res.DeletedCount, nil
}

func (r RepoMongoDB) update(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error) {
	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": socialDemGroup.ID},
		bson.M{"$set": socialDemGroupDocFromModel(socialDemGroup)},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return 0, errors.Wrap(err, "upsert social dem group")
	}

	return res.UpsertedCount + res.ModifiedCount, nil
}
