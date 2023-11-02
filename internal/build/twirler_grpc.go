package build

import (
	"context"
	twirler_v1 "github.com/Blancduman/banners-rotation/pkg/twirler/v1"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Blancduman/banners-rotation/internal/grpc/external"
)

func (b *Builder) ItemGRPCServer(ctx context.Context) (*grpc.Server, error) {
	bannerService, err := b.bannerService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "build banner service")
	}

	slotService, err := b.slotService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "build slot service")
	}

	socialDemGroupService, err := b.socialDemGroupService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "build social dem group service")
	}

	statService, err := b.statService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "build stat service")
	}

	s, err := b.GRPCServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "create gRPC server")
	}

	twirler_v1.RegisterTwirlerAPIServer(s, external.NewServer(external.Services{
		BannerService:         bannerService,
		SlotService:           slotService,
		SocialDemGroupService: socialDemGroupService,
		StatService:           statService,
	}, b.mongoCatalog.client))

	return s, err
}
