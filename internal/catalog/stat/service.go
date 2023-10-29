package stat

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

type Service struct {
	wg   sync.WaitGroup
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo} //nolint:exhaustruct
}

func (s *Service) Shutdown(_ context.Context) error {
	s.wg.Wait()

	return nil
}

func (s *Service) IncrementClickedCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	s.wg.Add(1)
	err := s.repo.IncrementClickedCount(ctx, slotID, bannerID, socialDemGroupID)
	s.wg.Done()

	return errors.Wrapf(err, "increment clicked count slot %d banner %d group %d", slotID, bannerID, socialDemGroupID)
}

func (s *Service) IncrementShownCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	s.wg.Add(1)
	err := s.repo.IncrementShownCount(ctx, slotID, bannerID, socialDemGroupID)
	s.wg.Done()

	return errors.Wrapf(err, "increment shown count slot %d banner %d group %d", slotID, bannerID, socialDemGroupID)
}
