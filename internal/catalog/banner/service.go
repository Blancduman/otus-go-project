package banner

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

func (s *Service) Get(ctx context.Context, id int64) (Banner, error) {
	banner, err := s.repo.Get(ctx, id)
	//if errors.Is(err, ErrNotFound) {
	//	return Banner{
	//		ID:          id,
	//		Description: "",
	//	}, nil
	//}

	return banner, errors.Wrapf(err, "get banner %d", id)
}

func (s *Service) Create(ctx context.Context, banner Banner) (int64, error) {
	if banner.Description == "" {
		return 0, errors.New("empty description")
	}

	s.wg.Add(1)
	ID, err := s.repo.Create(ctx, banner)
	s.wg.Done()

	return ID, errors.Wrap(err, "create banner")
}

func (s *Service) Update(ctx context.Context, banner Banner) error {
	s.wg.Add(1)
	_, err := s.repo.Update(ctx, banner)
	s.wg.Done()

	return errors.Wrap(err, "update banner")
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.wg.Add(1)
	_, err := s.repo.Delete(ctx, id)
	s.wg.Done()

	return errors.Wrapf(err, "delete banner %d", id)
}
