package socialdemgroup

import "context"

//go:generate mockery --name Repo
type Repo interface {
	Get(ctx context.Context, id int64) (SocialDemGroup, error)
	Create(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error)
	Update(ctx context.Context, socialDemGroup SocialDemGroup) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
