package auth

import (
	"testing"

	"github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"
	cacheMocks "github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/cache/mocks"
	repoMocks "github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/repository/mocks"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/jwt"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type testDeps struct {
	ctrl       *gomock.Controller
	userRepo   *repoMocks.MockUserRepository
	tokenCache *cacheMocks.MockTokenBlacklist
	svc        *auth
}

func setupTest(t *testing.T) *testDeps {
	t.Helper()

	setupEnv(t)

	ctrl := gomock.NewController(t)
	userRepo := repoMocks.NewMockUserRepository(ctrl)
	tokenCache := cacheMocks.NewMockTokenBlacklist(ctrl)
	svc := &auth{
		userRepo:   userRepo,
		tokenCache: tokenCache,
	}

	return &testDeps{
		ctrl:       ctrl,
		userRepo:   userRepo,
		tokenCache: tokenCache,
		svc:        svc,
	}
}

func setupEnv(t *testing.T) {
	t.Helper()
	t.Setenv("APP_SECRET_KEY", "test-secret-key-for-unit-tests")
	t.Setenv("APP_REFRESH_KEY", "test-refresh-key-for-unit-tests")
	viper.Set("APP_SECRET_KEY", "test-secret-key-for-unit-tests")
	viper.Set("APP_REFRESH_KEY", "test-refresh-key-for-unit-tests")
}

func generateTestRefreshToken(t *testing.T) string {
	t.Helper()
	token, err := jwt.GenerateToken(entity.UserCtx{Id: "user1", Email: "user@test.com"}, true)
	require.NoError(t, err)
	return token
}
