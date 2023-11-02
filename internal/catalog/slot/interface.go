package slot

import (
	"context"
)

//go:generate mockery --name Repo
type Repo interface {
	Get(ctx context.Context, id int64) (Slot, error)
	Create(ctx context.Context, slot Slot) (int64, error)
	Update(ctx context.Context, slot Slot) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	//AttachBanner(ctx context.Context, id int64, bannerID int64) (int64, error)
	//DetachBanner(ctx context.Context, id int64, bannerID int64) (int64, error)
}
