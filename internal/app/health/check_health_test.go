package health

import (
	"context"
	"errors"
	"testing"

	"go-simple-template/internal/core/domain/entity"
	cacheMocks "go-simple-template/internal/core/port/outbound/cache/mocks"
	repoMocks "go-simple-template/internal/core/port/outbound/repository/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testDeps struct {
	ctrl        *gomock.Controller
	healthRepo  *repoMocks.MockHealth
	healthCache *cacheMocks.MockHealth
	svc         *service
}

func setupTest(t *testing.T) *testDeps {
	t.Helper()

	ctrl := gomock.NewController(t)
	healthRepo := repoMocks.NewMockHealth(ctrl)
	healthCache := cacheMocks.NewMockHealth(ctrl)
	svc := &service{
		healthRepo:  healthRepo,
		healthCache: healthCache,
	}

	return &testDeps{
		ctrl:        ctrl,
		healthRepo:  healthRepo,
		healthCache: healthCache,
		svc:         svc,
	}
}

func TestCheckHealth(t *testing.T) {
	mongoErr := errors.New("mongo connection refused")
	redisErr := errors.New("redis connection refused")

	tests := []struct {
		name      string
		req       entity.CheckHealth
		ctx       func() context.Context
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.CheckHealth, err error)
	}{
		{
			name: "no checks requested",
			req:  entity.CheckHealth{MongoDB: false, Redis: false},
			mockSetup: func(deps *testDeps) {
				// no mock expectations
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.False(t, got.MongoDB)
				assert.False(t, got.Redis)
			},
		},
		{
			name: "mongodb only - success",
			req:  entity.CheckHealth{MongoDB: true, Redis: false},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(nil)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.True(t, got.MongoDB)
				assert.False(t, got.Redis)
			},
		},
		{
			name: "mongodb only - failure",
			req:  entity.CheckHealth{MongoDB: true, Redis: false},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(mongoErr)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
				assert.ErrorIs(t, err, mongoErr)
				assert.Nil(t, got)
			},
		},
		{
			name: "redis only - success",
			req:  entity.CheckHealth{MongoDB: false, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(nil)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.False(t, got.MongoDB)
				assert.True(t, got.Redis)
			},
		},
		{
			name: "redis only - failure",
			req:  entity.CheckHealth{MongoDB: false, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(redisErr)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
				assert.ErrorIs(t, err, redisErr)
				assert.Nil(t, got)
			},
		},
		{
			name: "both - all success",
			req:  entity.CheckHealth{MongoDB: true, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(nil)
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(nil)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.True(t, got.MongoDB)
				assert.True(t, got.Redis)
			},
		},
		{
			name: "both - mongodb fails",
			req:  entity.CheckHealth{MongoDB: true, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(mongoErr)
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(nil).AnyTimes()
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "both - redis fails",
			req:  entity.CheckHealth{MongoDB: true, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(nil).AnyTimes()
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(redisErr)
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "both - all fail",
			req:  entity.CheckHealth{MongoDB: true, Redis: true},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(mongoErr).AnyTimes()
				deps.healthCache.EXPECT().CheckHealth(gomock.Any()).Return(redisErr).AnyTimes()
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "context cancelled",
			req:  entity.CheckHealth{MongoDB: true, Redis: false},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			mockSetup: func(deps *testDeps) {
				deps.healthRepo.EXPECT().CheckHealth(gomock.Any()).Return(context.Canceled).AnyTimes()
			},
			assertFn: func(t *testing.T, got *entity.CheckHealth, err error) {
				assert.Error(t, err)
				assert.ErrorIs(t, err, context.Canceled)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			got, err := deps.svc.CheckHealth(ctx, tt.req)
			tt.assertFn(t, got, err)
		})
	}
}
