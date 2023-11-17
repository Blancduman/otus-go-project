package slot

import (
	"context"
	"sync"

	"github.com/pkg/errors"
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

func (s *Service) Create(ctx context.Context, slot Slot) (int64, error) {
	if slot.Description == "" {
		return 0, errors.New("empty description")
	}

	s.wg.Add(1)
	ID, err := s.repo.Create(ctx, slot)
	s.wg.Done()

	return ID, errors.Wrap(err, "create slot")
}

func (s *Service) Get(ctx context.Context, id int64) (Slot, error) {
	slot, err := s.repo.Get(ctx, id)

	return slot, errors.Wrapf(err, "get slot %d", id)
}

func (s *Service) Update(ctx context.Context, slot Slot) error {
	s.wg.Add(1)
	_, err := s.repo.Update(ctx, slot)
	s.wg.Done()

	return errors.Wrap(err, "update slot")
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.wg.Add(1)
	_, err := s.repo.Delete(ctx, id)
	s.wg.Done()

	return errors.Wrapf(err, "delete slot %d", id)
}
