package stat_test

import (
	"context"
	"testing"

	"github.com/Blancduman/banners-rotation/internal/catalog/stat"
	"github.com/Blancduman/banners-rotation/internal/catalog/stat/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_GetStat(t *testing.T) {
	t.Parallel()

	defaultSlotStat := stat.SlotStat{}

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlotStat, errors.New("foo")).Once()
		s1, err1 := service.GetStat(ctx, 0)
		require.Error(t, err1)
		require.Equal(t, defaultSlotStat, s1)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlotStat, stat.ErrNotFound).Once()
		s2, err2 := service.GetStat(ctx, 0)
		require.Error(t, err2)
		require.Equal(t, defaultSlotStat, s2)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID:     1,
			BannerStat: map[int64]stat.SocialDemGroupStat{},
		}, nil).Once()
		s1, err1 := service.GetStat(ctx, 1)
		require.NoError(t, err1)
		require.Equal(t, stat.SlotStat{
			SlotID:     1,
			BannerStat: map[int64]stat.SocialDemGroupStat{},
		}, s1)
	})
}

func Test_IncrementClickedCount(t *testing.T) { //nolint: dupl
	t.Parallel()

	defaultSlotStat := stat.SlotStat{}

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlotStat, errors.New("foo")).Once()
		err1 := service.IncrementClickedCount(ctx, 1, 1, 1)
		require.Error(t, err1)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				2: map[int64]stat.Stat{},
			},
		}, nil).Once()
		err11 := service.IncrementClickedCount(ctx, 1, 1, 1)
		require.Error(t, err11)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementClickedCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(errors.New("foo")).Once()
		err2 := service.IncrementClickedCount(ctx, 1, 1, 1)
		require.Error(t, err2)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementClickedCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(nil).Once()
		p.On(
			"Produce",
			mock.Anything,
			mock.AnythingOfType("payload.ClickStat"),
		).Return(errors.New("foo")).Once()
		err3 := service.IncrementClickedCount(ctx, 1, 1, 1)
		require.Error(t, err3)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementClickedCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(nil).Once()
		p.On(
			"Produce",
			mock.Anything,
			mock.AnythingOfType("payload.ClickStat"),
		).Return(nil).Once()
		err2 := service.IncrementClickedCount(ctx, 1, 1, 1)
		require.NoError(t, err2)
	})
}

func Test_IncrementShownCount(t *testing.T) { //nolint: dupl
	t.Parallel()

	defaultSlotStat := stat.SlotStat{}

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlotStat, errors.New("foo")).Once()
		err1 := service.IncrementShownCount(ctx, 1, 1, 1)
		require.Error(t, err1)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				2: map[int64]stat.Stat{},
			},
		}, nil).Once()
		err11 := service.IncrementShownCount(ctx, 1, 1, 1)
		require.Error(t, err11)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementShownCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(errors.New("foo")).Once()
		err2 := service.IncrementShownCount(ctx, 1, 1, 1)
		require.Error(t, err2)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementShownCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(nil).Once()
		p.On(
			"Produce",
			mock.Anything,
			mock.AnythingOfType("payload.ShownStat"),
		).Return(errors.New("foo")).Once()
		err3 := service.IncrementShownCount(ctx, 1, 1, 1)
		require.Error(t, err3)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On("GetStat", mock.Anything, mock.AnythingOfType("int64")).Return(stat.SlotStat{
			SlotID: 1,
			BannerStat: map[int64]stat.SocialDemGroupStat{
				1: map[int64]stat.Stat{},
			},
		}, nil).Once()
		m.On(
			"IncrementShownCount",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(nil).Once()
		p.On(
			"Produce",
			mock.Anything,
			mock.AnythingOfType("payload.ShownStat"),
		).Return(nil).Once()
		err1 := service.IncrementShownCount(ctx, 1, 1, 1)
		require.NoError(t, err1)
	})
}

func Test_AddBannerToSlot(t *testing.T) {
	t.Parallel()

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On(
			"AddBannerToSlot",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("[]int64"),
		).Return(errors.New("foo")).Once()
		err := service.AddBannerToSlot(ctx, 1, 2, []int64{3})
		require.Error(t, err)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On(
			"AddBannerToSlot",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("[]int64"),
		).Return(nil).Once()
		err := service.AddBannerToSlot(ctx, 1, 2, []int64{3})
		require.NoError(t, err)
	})
}

func Test_RemoveBannerFromSlot(t *testing.T) {
	t.Parallel()

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On(
			"RemoveBannerFromSlot",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(errors.New("foo")).Once()
		err := service.RemoveBannerFromSlot(ctx, 1, 2)
		require.Error(t, err)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		p := mocks.NewProducer(t)
		service := stat.NewService(m, p)

		m.On(
			"RemoveBannerFromSlot",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("int64"),
		).Return(nil).Once()
		err := service.RemoveBannerFromSlot(ctx, 1, 2)
		require.NoError(t, err)
	})
}
