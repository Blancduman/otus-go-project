package socialdemgroup_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Blancduman/banners-rotation/internal/catalog/socialdemgroup"
	"github.com/Blancduman/banners-rotation/internal/catalog/socialdemgroup/mocks"
)

func Test_Get(t *testing.T) {
	t.Parallel()

	defaultGroup := socialdemgroup.SocialDemGroup{}

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := socialdemgroup.NewService(m)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultGroup, errors.New("foo")).Once()
		s1, err1 := service.Get(ctx, 0)
		require.Error(t, err1)
		require.Equal(t, defaultGroup, s1)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultGroup, socialdemgroup.ErrNotFound).Once()
		s2, err2 := service.Get(ctx, 1)
		require.Error(t, err2)
		require.Equal(t, defaultGroup, s2)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := socialdemgroup.NewService(m)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(socialdemgroup.SocialDemGroup{
			ID:          1,
			Description: "Test 1",
		}, nil).Once()
		b, err := service.Get(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, socialdemgroup.SocialDemGroup{
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
		service := socialdemgroup.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("socialdemgroup.SocialDemGroup")).Return(int64(0), errors.New("foo")).Once()
		id, err := service.Create(ctx, socialdemgroup.Fixture1())
		require.Error(t, err)
		require.Equal(t, int64(0), id)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := socialdemgroup.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("socialdemgroup.SocialDemGroup")).Return(int64(1), nil).Once()
		id, err := service.Create(ctx, socialdemgroup.Fixture3())
		require.NoError(t, err)
		require.Equal(t, int64(1), id)
	})
}
