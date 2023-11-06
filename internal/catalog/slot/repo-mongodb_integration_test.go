//go:build integration

package slot

import (
	"context"
	"fmt"
	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

type MongoDBRepoTestSuite struct {
	suite.Suite
	repo *RepoMongoDB
	conn *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

func (s *MongoDBRepoTestSuite) SetupSuite() {
	ctx := context.TODO()

	cfg, err := config.Load()
	s.Require().NoError(err)

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.DSN))
	s.Require().NoError(err)

	s.conn = conn
	s.db = conn.Database(fmt.Sprintf("test-%s", uuid.New()))
	s.coll = s.db.Collection("slot", options.Collection())
	s.repo = NewMongoDBRepo(s.coll)
}

func (s *MongoDBRepoTestSuite) SetupTest() {
	docs := make([]interface{}, 3)

	for i, m := range []Slot{Fixture1(), Fixture2(), Fixture3()} {
		docs[i] = slotDocFromModel(m)
	}

	_, err := s.coll.InsertMany(context.TODO(), docs)
	s.Require().NoError(err)
}

func (s *MongoDBRepoTestSuite) TearDownSuite() {
	ctx := context.TODO()
	err := s.db.Drop(ctx)
	s.Require().NoError(err)
	err = s.conn.Disconnect(ctx)
	s.Require().NoError(err)
}

func (s *MongoDBRepoTestSuite) TearDownTest() {
	err := s.coll.Drop(context.TODO())
	s.Require().NoError(err)
}

func TestRepoMongodbTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MongoDBRepoTestSuite))
}

func (s *MongoDBRepoTestSuite) Test_Get() {
	s.Run("Founded", func() {
		ctx := context.TODO()
		want := Fixture1()
		got, err := s.repo.Get(ctx, want.ID)
		s.Require().NoError(err)
		s.Require().Equal(want, got)
	})

	s.Run("Not founded", func() {
		ctx := context.TODO()
		_, err := s.repo.Get(ctx, 4)
		s.Require().Error(err)
	})
}

func (s *MongoDBRepoTestSuite) Test_Create() {
	want := int64(4)

	id, err := s.repo.Create(context.TODO(), Slot{
		ID:          0,
		Description: "Test 4",
	})
	s.Require().NoError(err)
	s.Require().Equal(want, id)
}
