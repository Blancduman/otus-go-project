package stat

import "context"

//go:generate mockery --name Repo
type Repo interface {
	IncrementClickedCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error
	IncrementShownCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error
	//AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error
	//RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error
}
