package external

import (
	"context"

	"github.com/Blancduman/banners-rotation/internal/catalog/banner"
	"github.com/Blancduman/banners-rotation/internal/catalog/slot"
	"github.com/Blancduman/banners-rotation/internal/catalog/socialdemgroup"
	"github.com/Blancduman/banners-rotation/internal/catalog/stat"
	"github.com/Blancduman/banners-rotation/internal/ucb1"
	twirler_v1 "github.com/Blancduman/banners-rotation/pkg/twirler/v1"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	services Services
	client   *mongo.Client
	twirler_v1.UnimplementedTwirlerAPIServer
}

type Services struct {
	BannerService         BannerService
	SlotService           SlotService
	SocialDemGroupService SocialDemGroupService
	StatService           StatService
}

type BannerService interface {
	Create(context.Context, banner.Banner) (int64, error)
}

type SlotService interface {
	Create(ctx context.Context, slot slot.Slot) (int64, error)
}

type SocialDemGroupService interface {
	Get(ctx context.Context, id int64) (socialdemgroup.SocialDemGroup, error)
	Create(ctx context.Context, socialDemGroup socialdemgroup.SocialDemGroup) (int64, error)
	GetAll(ctx context.Context) ([]socialdemgroup.SocialDemGroup, error)
}

type StatService interface {
	GetStat(ctx context.Context, slotID int64) (stat.SlotStat, error)
	IncrementClickedCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error
	IncrementShownCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error
	AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error
	RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error
}

func NewServer(s Services, c *mongo.Client) *Server {
	return &Server{
		services: s,
		client:   c,
	}
}

func (s *Server) AttachBanner(
	ctx context.Context,
	req *twirler_v1.AttachBannerToSlotRequest,
) (*twirler_v1.AttachBannerToSlotResponse, error) {
	groups, err := s.services.SocialDemGroupService.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get all social dem groups from social dem group service")
	}

	groupIDs := make([]int64, len(groups))

	for i, v := range groups {
		groupIDs[i] = v.ID
	}

	err = s.services.StatService.AddBannerToSlot(ctx, int64(req.GetSlotID()), int64(req.GetBannerID()), groupIDs)
	if err != nil {
		return nil, errors.Wrap(err, "add banner to slot stat service")
	}

	return &twirler_v1.AttachBannerToSlotResponse{}, nil
}

func (s *Server) DetachBanner(
	ctx context.Context,
	req *twirler_v1.DetachBannerToSlotRequest,
) (*twirler_v1.DetachBannerToSlotResponse, error) {
	err := s.services.StatService.RemoveBannerFromSlot(ctx, int64(req.GetSlotID()), int64(req.GetBannerID()))
	if err != nil {
		return nil, errors.Wrap(err, "detach banner from service")
	}

	return &twirler_v1.DetachBannerToSlotResponse{}, nil
}

func (s *Server) IncrementCount(
	ctx context.Context,
	req *twirler_v1.IncrementCountRequest,
) (*twirler_v1.IncrementCountResponse, error) {
	err := s.services.StatService.IncrementClickedCount(
		ctx,
		int64(req.GetSlotID()),
		int64(req.GetBannerID()),
		int64(req.GetBannerID()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "increment count from service")
	}

	return &twirler_v1.IncrementCountResponse{}, nil
}

func (s *Server) Gimme(ctx context.Context, req *twirler_v1.GimmeRequest) (*twirler_v1.GimmeResponse, error) {
	slotStat, err := s.services.StatService.GetStat(ctx, int64(req.GetSlotID()))
	if err != nil {
		return nil, errors.Wrap(err, "get stat from service")
	}

	if len(slotStat.BannerStat) == 0 {
		return nil, errors.New("no banner attached to slot")
	}

	sdGroup, err := s.services.SocialDemGroupService.Get(ctx, int64(req.GetSocialGroupId()))
	if err != nil {
		return nil, errors.Wrap(err, "get social dem group from service")
	}

	stats := slotStat.BannerStat.ToUCB1()

	bannerID := ucb1.Next(stats, sdGroup.ID)

	err = s.services.StatService.IncrementShownCount(ctx, slotStat.SlotID, bannerID, sdGroup.ID)
	if err != nil {
		return nil, errors.Wrap(err, "can not increment shown count")
	}

	return &twirler_v1.GimmeResponse{BannerID: uint64(bannerID)}, nil
}

func (s *Server) CreateBanner(
	ctx context.Context,
	req *twirler_v1.BannerCreateRequest,
) (*twirler_v1.BannerCreateResponse, error) {
	ID, err := s.services.BannerService.Create(ctx, banner.Banner{
		ID:          0,
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create banner from service")
	}

	return &twirler_v1.BannerCreateResponse{ID: ID}, nil
}

func (s *Server) CreateSlot(
	ctx context.Context,
	req *twirler_v1.SlotCreateRequest,
) (*twirler_v1.SlotCreateResponse, error) {
	ID, err := s.services.SlotService.Create(ctx, slot.Slot{
		ID:          0,
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create slot from service")
	}

	return &twirler_v1.SlotCreateResponse{ID: ID}, nil
}

func (s *Server) CreateSocialDemGroup(
	ctx context.Context,
	req *twirler_v1.SocialDemGroupCreateRequest,
) (*twirler_v1.SocialDemGroupCreateResponse, error) {
	ID, err := s.services.SocialDemGroupService.Create(ctx, socialdemgroup.SocialDemGroup{
		ID:          0,
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create social dem group from service")
	}

	return &twirler_v1.SocialDemGroupCreateResponse{ID: ID}, nil
}
