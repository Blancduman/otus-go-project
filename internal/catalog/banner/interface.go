package banner

import (
	"context"
)

//go:generate mockery --name Repo
type Repo interface {
	Get(ctx context.Context, id int64) (Banner, error)
	Create(ctx context.Context, banner Banner) (int64, error)
	Update(ctx context.Context, banner Banner) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
