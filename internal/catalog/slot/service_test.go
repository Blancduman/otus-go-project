package slot_test

import (
	"context"
	"testing"

	"github.com/Blancduman/banners-rotation/internal/catalog/slot"
	"github.com/Blancduman/banners-rotation/internal/catalog/slot/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	t.Parallel()

	defaultSlot := slot.Slot{}

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := slot.NewService(m)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlot, errors.New("foo")).Once()
		s1, err1 := service.Get(ctx, 0)
		require.Error(t, err1)
		require.Equal(t, defaultSlot, s1)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(defaultSlot, slot.ErrNotFound).Once()
		s2, err2 := service.Get(ctx, 1)
		require.Error(t, err2)
		require.Equal(t, defaultSlot, s2)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := slot.NewService(m)

		m.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(slot.Slot{
			ID:          1,
			Description: "Test 1",
		}, nil).Once()
		b, err := service.Get(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, slot.Slot{
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
		service := slot.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("Slot")).Return(int64(0), errors.New("foo")).Once()
		id, err := service.Create(ctx, slot.Fixture1())
		require.Error(t, err)
		require.Equal(t, int64(0), id)
	})

	t.Run("repo fine", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()
		m := mocks.NewRepo(t)
		service := slot.NewService(m)

		m.On("Create", mock.Anything, mock.AnythingOfType("Slot")).Return(int64(1), nil).Once()
		id, err := service.Create(ctx, slot.Fixture3())
		require.NoError(t, err)
		require.Equal(t, int64(1), id)
	})
}
