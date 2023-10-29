package banner

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type RepoMongoDB struct {
	collection *mongo.Collection
}

func NewMongoDBRepo(collection *mongo.Collection) *RepoMongoDB {
	return &RepoMongoDB{
		collection: collection,
	}
}

func (r RepoMongoDB) Get(ctx context.Context, id int64) (Banner, error) {
	var doc bannerDoc

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return Banner{}, ErrNotFound
	}

	return doc.toModel(), errors.Wrap(err, "find banner by id")
}

func (r RepoMongoDB) Create(ctx context.Context, banner Banner) (int64, error) {
	session, err := r.collection.Database().Client().StartSession(options.Session())
	if err != nil {
		return 0, errors.Wrap(err, "start session")
	}

	defer session.EndSession(ctx)

	res, err := session.WithTransaction(ctx, func(ctxSession mongo.SessionContext) (interface{}, error) {
		var doc bannerDoc
		err := r.collection.FindOne(ctxSession, bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&doc)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return int32(0), errors.Wrap(err, "get last document")
		}

		lastID := doc.ID
		banner.ID = lastID + 1

		_, err = r.collection.InsertOne(ctxSession, bannerDocFromModel(banner))
		if err != nil {
			return int32(0), errors.Wrap(err, "insert document")
		}

		return banner.ID, nil
	}, options.Transaction().SetReadPreference(readpref.Primary()))
	if err != nil {
		return 0, errors.Wrap(err, "exec transaction")
	}

	if id, ok := res.(int64); ok {
		return id, nil
	}

	return 0, ErrInvalidBannerIDCreated
}

func (r RepoMongoDB) Update(ctx context.Context, banner Banner) (int64, error) {
	res, err := r.update(ctx, banner)

	return res, errors.Wrap(err, "update banner")
}

func (r RepoMongoDB) Delete(ctx context.Context, id int64) (int64, error) {
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, errors.Wrap(err, "delete banner by id")
	}

	return res.DeletedCount, nil
}

func (r RepoMongoDB) update(ctx context.Context, banner Banner) (int64, error) {
	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": banner.ID},
		bson.M{"$set": bannerDocFromModel(banner)},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return 0, errors.Wrap(err, "upsert banner")
	}

	return res.UpsertedCount + res.ModifiedCount, nil
}
