package banner_test

import (
	"context"
	"testing"

	"github.com/Blancduman/banners-rotation/internal/catalog/banner"
	"github.com/Blancduman/banners-rotation/internal/catalog/banner/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	t.Parallel()

	defaultBanner := banner.Banner{} //nolint: exhaustruct

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := banner.NewService(m)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultBanner, errors.New("foo")).Once()
		b1, err1 := service.Get(ctx, 0)
		require.Error(t, err1)
		require.Equal(t, defaultBanner, b1)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultBanner, banner.ErrNotFound).Once()
		b2, err2 := service.Get(ctx, 1)
		require.Error(t, err2)
		require.Equal(t, defaultBanner, b2)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(banner.Banner{
			ID:          1,
			Description: "Test 1",
		}, nil).Once()
		service := banner.NewService(m)
		b, err := service.Get(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, banner.Banner{
			ID:          1,
			Description: "Test 1",
		}, b)
	})
}

func Test_Create(t *testing.T) {
	t.Parallel()

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := banner.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("Banner")).Return(int64(0), errors.New("foo")).Once()
		id, err := service.Create(ctx, banner.Fixture1())
		require.Error(t, err)
		require.Equal(t, int64(0), id)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := banner.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("Banner")).Return(int64(1), nil).Once()
		id, err := service.Create(ctx, banner.Fixture3())
		require.NoError(t, err)
		require.Equal(t, int64(1), id)
	})
}
