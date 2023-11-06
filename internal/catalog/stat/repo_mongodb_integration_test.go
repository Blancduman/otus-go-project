//go:build integration

package stat

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
	s.coll = s.db.Collection("stat", options.Collection())
	s.repo = NewMongoDBRepo(s.coll)
}

func (s *MongoDBRepoTestSuite) SetupTest() {
	docs := []interface{}{fixture1(), fixture2(), fixture3()}

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

func (s *MongoDBRepoTestSuite) Test_GetStat() {
	s.Run("Found", func() {
		ctx := context.TODO()
		want := fixture1().toModel()
		got, err := s.repo.GetStat(ctx, want.SlotID)
		s.Require().NoError(err)
		s.Require().Equal(want, got)
	})

	s.Run("Not found", func() {
		ctx := context.TODO()
		_, err := s.repo.GetStat(ctx, 10)
		s.Require().Error(err)
	})
}

func (s *MongoDBRepoTestSuite) Test_IncrementClickedCount() {
	var bannerID int64
	var socialGroupID int64

	ctx := context.TODO()
	slSt := fixture1()

	for k, v := range slSt.BannerStat {
		bannerID = k
		for k2, _ := range v {
			socialGroupID = k2

			break
		}

		break
	}

	err := s.repo.IncrementClickedCount(ctx, slSt.ID, bannerID, socialGroupID)
	s.Require().NoError(err)

	slotStat, err := s.repo.GetStat(ctx, slSt.ID)
	s.Require().NoError(err)
	s.Require().Equal(int64(2), slotStat.BannerStat[bannerID][socialGroupID].Clicked)
}

func (s *MongoDBRepoTestSuite) Test_IncrementShownCount() {
	var bannerID int64
	var socialGroupID int64

	ctx := context.TODO()
	slSt := fixture1()

	for k, v := range slSt.BannerStat {
		bannerID = k
		for k2, _ := range v {
			socialGroupID = k2

			break
		}

		break
	}

	err := s.repo.IncrementShownCount(ctx, slSt.ID, bannerID, socialGroupID)
	s.Require().NoError(err)

	slotStat, err := s.repo.GetStat(ctx, slSt.ID)
	s.Require().NoError(err)
	s.Require().Equal(int64(2), slotStat.BannerStat[bannerID][socialGroupID].Shown)
}

func (s *MongoDBRepoTestSuite) Test_AddBannerToSlot() {
	ctx := context.TODO()
	slSt := fixture1()

	err := s.repo.AddBannerToSlot(ctx, slSt.ID, 20, []int64{1, 2, 3, 4, 5, 6})
	s.Require().NoError(err)

	slotStat, err := s.repo.GetStat(ctx, slSt.ID)
	s.Require().NoError(err)
	s.Require().NotNil(slotStat.BannerStat[20])

	for k, _ := range slotStat.BannerStat[20] {
		s.Require().Contains([]int64{1, 2, 3, 4, 5, 6}, k)
	}
}

func (s *MongoDBRepoTestSuite) Test_RemoveBannerFromSlot() {
	var bannerID int64

	ctx := context.TODO()
	slSt := fixture1()

	for k, _ := range slSt.BannerStat {
		bannerID = k

		break
	}

	err := s.repo.RemoveBannerFromSlot(ctx, slSt.ID, bannerID)
	s.Require().NoError(err)

	slotStat, err := s.repo.GetStat(ctx, slSt.ID)
	s.Require().NoError(err)
	s.Require().Nil(slotStat.BannerStat[slSt.ID])
}
