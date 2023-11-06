// Code generated by mockery v2.34.0. DO NOT EDIT.

package mocks

import (
	context "context"

	stat "github.com/Blancduman/banners-rotation/internal/catalog/stat"
	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// AddBannerToSlot provides a mock function with given fields: ctx, slotID, bannerID, socialDemGroupIDs
func (_m *Repo) AddBannerToSlot(ctx context.Context, slotID int64, bannerID int64, socialDemGroupIDs []int64) error {
	ret := _m.Called(ctx, slotID, bannerID, socialDemGroupIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, []int64) error); ok {
		r0 = rf(ctx, slotID, bannerID, socialDemGroupIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStat provides a mock function with given fields: ctx, slotID
func (_m *Repo) GetStat(ctx context.Context, slotID int64) (stat.SlotStat, error) {
	ret := _m.Called(ctx, slotID)

	var r0 stat.SlotStat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (stat.SlotStat, error)); ok {
		return rf(ctx, slotID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) stat.SlotStat); ok {
		r0 = rf(ctx, slotID)
	} else {
		r0 = ret.Get(0).(stat.SlotStat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, slotID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementClickedCount provides a mock function with given fields: ctx, slotID, bannerID, socialDemGroupID
func (_m *Repo) IncrementClickedCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	ret := _m.Called(ctx, slotID, bannerID, socialDemGroupID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64) error); ok {
		r0 = rf(ctx, slotID, bannerID, socialDemGroupID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IncrementShownCount provides a mock function with given fields: ctx, slotID, bannerID, socialDemGroupID
func (_m *Repo) IncrementShownCount(ctx context.Context, slotID int64, bannerID int64, socialDemGroupID int64) error {
	ret := _m.Called(ctx, slotID, bannerID, socialDemGroupID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64) error); ok {
		r0 = rf(ctx, slotID, bannerID, socialDemGroupID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveBannerFromSlot provides a mock function with given fields: ctx, slotID, bannerID
func (_m *Repo) RemoveBannerFromSlot(ctx context.Context, slotID int64, bannerID int64) error {
	ret := _m.Called(ctx, slotID, bannerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, slotID, bannerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}