package health

import (
	"testing"

	cacheMocks "github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/cache/mocks"
	repoMocks "github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/repository/mocks"

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
