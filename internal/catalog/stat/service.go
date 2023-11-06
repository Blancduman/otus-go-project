package stat

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Blancduman/banners-rotation/internal/reporter/payload"
	"github.com/pkg/errors"
)

type Service struct {
	wg       sync.WaitGroup
	repo     Repo
	producer Producer
}

//go:generate mockery --filename producer.go --name Producer
type Producer interface {
	Produce(ctx context.Context, payload payload.Payload) error
}

func NewService(repo Repo, producer Producer) *Service {
	return &Service{repo: repo, producer: producer} //nolint:exhaustruct
}

func (s *Service) Shutdown(_ context.Context) error {
	s.wg.Wait()

	return nil
}

func (s *Service) GetStat(ctx context.Context, slotID int64) (SlotStat, error) {
	stat, err := s.repo.GetStat(ctx, slotID)

	return stat, errors.Wrapf(err, "get stat %d", slotID)
}

func (s *Service) IncrementClickedCount( //nolint: dupl
	ctx context.Context,
	slotID int64,
	bannerID int64,
	socialDemGroupID int64,
) error {
	slotStat, err := s.repo.GetStat(ctx, slotID)
	if err != nil {
		return errors.Wrap(err, "get slot from service")
	}

	if _, ok := slotStat.BannerStat[bannerID]; !ok {
		return fmt.Errorf("banner %d is not attached to slot %d", slotID, bannerID)
	}

	s.wg.Add(1)
	defer s.wg.Done()

	err = s.repo.IncrementClickedCount(ctx, slotID, bannerID, socialDemGroupID)
	if err != nil {
		return errors.Wrapf(err, "increment clicked count slot %d banner %d group %d", slotID, bannerID, socialDemGroupID)
	}

	err = s.producer.Produce(ctx, payload.ClickStat{
		SlotID:           slotID,
		BannerID:         bannerID,
		SocialDemGroupID: socialDemGroupID,
		Timestamp:        time.Now(),
	})
	if err != nil {
		return errors.Wrapf(
			err,
			"faild to send click message to kafka slot %d banner %d group %d",
			slotID,
			bannerID,
			socialDemGroupID,
		)
	}

	return nil
}

func (s *Service) IncrementShownCount( //nolint: dupl
	ctx context.Context,
	slotID int64,
	bannerID int64,
	socialDemGroupID int64,
) error {
	slotStat, err := s.repo.GetStat(ctx, slotID)
	if err != nil {
		return errors.Wrap(err, "get slot from service")
	}

	if _, ok := slotStat.BannerStat[bannerID]; !ok {
		return fmt.Errorf("banner %d is not attached to slot %d", slotID, bannerID)
	}

	s.wg.Add(1)
	defer s.wg.Done()

	err = s.repo.IncrementShownCount(ctx, slotID, bannerID, socialDemGroupID)
	if err != nil {
		return errors.Wrapf(err, "increment shown count slot %d banner %d group %d", slotID, bannerID, socialDemGroupID)
	}

	err = s.producer.Produce(ctx, payload.ShownStat{
		SlotID:           slotID,
		BannerID:         bannerID,
		SocialDemGroupID: socialDemGroupID,
		Timestamp:        time.Now(),
	})
	if err != nil {
		return errors.Wrapf(
			err,
			"faild to send shown message to kafka slot %d banner %d group %d",
			slotID,
			bannerID,
			socialDemGroupID,
		)
	}

	return nil
}

func (s *Service) AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error {
	s.wg.Add(1)
	err := s.repo.AddBannerToSlot(ctx, slotID, bannerID, socialDemGroupIDs)
	s.wg.Done()

	return errors.Wrapf(err, "add banner from slot %d banner %d", slotID, bannerID)
}

func (s *Service) RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error {
	s.wg.Add(1)
	err := s.repo.RemoveBannerFromSlot(ctx, slotID, bannerID)
	s.wg.Done()

	return errors.Wrapf(err, "remove banner from slot %d banner %d", slotID, bannerID)
}
