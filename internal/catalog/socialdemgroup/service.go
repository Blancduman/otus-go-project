package socialdemgroup

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

func (s *Service) Get(ctx context.Context, id int64) (SocialDemGroup, error) {
	socialDemGroup, err := s.repo.Get(ctx, id)
	//if errors.Is(err, ErrNotFound) {
	//	return SocialDemGroup{
	//		ID:          id,
	//		Description: "",
	//	}, nil
	//}

	return socialDemGroup, errors.Wrapf(err, "get social dem group %d", id)
}

func (s *Service) GetAll(ctx context.Context) ([]SocialDemGroup, error) {
	socialDemGroups, err := s.repo.GetAll(ctx)
	if errors.Is(err, ErrNotFound) {
		return nil, nil
	}

	return socialDemGroups, errors.Wrap(err, "get all social dem group")
}

func (s *Service) Create(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error) {
	if socialDemGroup.Description == "" {
		return 0, errors.New("empty description")
	}

	s.wg.Add(1)
	ID, err := s.repo.Create(ctx, socialDemGroup)
	s.wg.Done()

	return ID, errors.Wrap(err, "create social dem group")
}

func (s *Service) Update(ctx context.Context, socialDemGroup SocialDemGroup) error {
	s.wg.Add(1)
	_, err := s.repo.Update(ctx, socialDemGroup)
	s.wg.Done()

	return errors.Wrap(err, "update social dem group")
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	s.wg.Add(1)
	_, err := s.repo.Delete(ctx, id)
	s.wg.Done()

	return errors.Wrapf(err, "delete social dem group %d", id)
}
