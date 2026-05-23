package auth

import (
	"context"
	"testing"

	"go-simple-template/internal/core/domain/entity"
	cacheMocks "go-simple-template/internal/core/port/outbound/cache/mocks"
	repoMocks "go-simple-template/internal/core/port/outbound/repository/mocks"
	"go-simple-template/internal/pkg/crypto"
	"go-simple-template/internal/pkg/errs"
	"go-simple-template/internal/pkg/jwt"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type testDeps struct {
	ctrl       *gomock.Controller
	userRepo   *repoMocks.MockUserRepository
	tokenCache *cacheMocks.MockToken
	svc        *auth
}

func setupTest(t *testing.T) *testDeps {
	t.Helper()

	// Set env vars needed by jwt package
	t.Setenv("APP_SECRET_KEY", "test-secret-key-for-unit-tests")
	t.Setenv("APP_REFRESH_KEY", "test-refresh-key-for-unit-tests")

	// Set viper keys directly as well for fast reading in tests
	viper.Set("APP_SECRET_KEY", "test-secret-key-for-unit-tests")
	viper.Set("APP_REFRESH_KEY", "test-refresh-key-for-unit-tests")

	ctrl := gomock.NewController(t)
	userRepo := repoMocks.NewMockUserRepository(ctrl)
	tokenCache := cacheMocks.NewMockToken(ctrl)
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

// --- Register Tests ---

func TestRegister(t *testing.T) {
	tests := []struct {
		name      string
		req       entity.User
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.User, err error)
	}{
		{
			name: "success",
			req:  entity.User{Email: "new@test.com", Password: "password123", Name: "Test User"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewNotFound("user not found"))
				deps.userRepo.EXPECT().
					Insert(gomock.Any(), gomock.Any()).
					Return(&entity.User{
						Base:  entity.Base{Id: "abc123"},
						Name:  "Test User",
						Email: "new@test.com",
					}, nil)
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				require.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, "abc123", got.Id)
				assert.Equal(t, "new@test.com", got.Email)
			},
		},
		{
			name: "user already exists",
			req:  entity.User{Email: "existing@test.com", Password: "password123"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(&entity.User{Email: "existing@test.com"}, nil)
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrConflict, errs.GetCode(err))
			},
		},
		{
			name: "find user returns unexpected error",
			req:  entity.User{Email: "test@test.com", Password: "password123"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewInternal(nil, "db connection failed"))
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrInternal, errs.GetCode(err))
			},
		},
		{
			name: "insert fails",
			req:  entity.User{Email: "new@test.com", Password: "password123"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewNotFound("user not found"))
				deps.userRepo.EXPECT().
					Insert(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewInternal(nil, "insert failed"))
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			got, err := deps.svc.Register(context.Background(), tt.req)
			tt.assertFn(t, got, err)
		})
	}
}

// --- Login Tests ---

func TestLogin(t *testing.T) {
	hashedPw, _ := crypto.HashPassword("correctpassword", 4)

	tests := []struct {
		name      string
		req       entity.User
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.AuthToken, err error)
	}{
		{
			name: "success",
			req:  entity.User{Email: "user@test.com", Password: "correctpassword"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(&entity.User{
						Base:     entity.Base{Id: "user1"},
						Email:    "user@test.com",
						Password: hashedPw,
					}, nil)
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				require.NoError(t, err)
				assert.NotNil(t, got)
				assert.NotEmpty(t, got.AccessToken)
				assert.NotEmpty(t, got.RefreshToken)
			},
		},
		{
			name: "user not found returns unauthorized",
			req:  entity.User{Email: "nobody@test.com", Password: "password"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewNotFound("user not found"))
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "wrong password",
			req:  entity.User{Email: "user@test.com", Password: "wrongpassword"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(&entity.User{
						Base:     entity.Base{Id: "user1"},
						Email:    "user@test.com",
						Password: hashedPw,
					}, nil)
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "repo internal error",
			req:  entity.User{Email: "user@test.com", Password: "password"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewInternal(nil, "db error"))
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrInternal, errs.GetCode(err))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			got, err := deps.svc.Login(context.Background(), tt.req)
			tt.assertFn(t, got, err)
		})
	}
}

// --- Profile Tests ---

func TestProfile(t *testing.T) {
	tests := []struct {
		name      string
		req       entity.AuthToken
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.User, err error)
	}{
		{
			name: "success",
			req:  entity.AuthToken{Id: "user1"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(&entity.User{
						Base:  entity.Base{Id: "user1"},
						Name:  "Test User",
						Email: "user@test.com",
					}, nil)
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				require.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, "user1", got.Id)
				assert.Equal(t, "Test User", got.Name)
			},
		},
		{
			name: "user not found returns unauthorized",
			req:  entity.AuthToken{Id: "nonexistent"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewNotFound("user not found"))
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "repo internal error",
			req:  entity.AuthToken{Id: "user1"},
			mockSetup: func(deps *testDeps) {
				deps.userRepo.EXPECT().
					FindOne(gomock.Any(), gomock.Any()).
					Return(nil, errs.NewInternal(nil, "db error"))
			},
			assertFn: func(t *testing.T, got *entity.User, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrInternal, errs.GetCode(err))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			got, err := deps.svc.Profile(context.Background(), tt.req)
			tt.assertFn(t, got, err)
		})
	}
}

// --- Logout Tests ---

func TestLogout(t *testing.T) {
	tests := []struct {
		name      string
		reqFn     func(t *testing.T, deps *testDeps) entity.AuthToken
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, err error)
	}{
		{
			name: "success",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: generateTestRefreshToken(t)}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					Blacklist(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			assertFn: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "invalid refresh token",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: "invalid-token"}
			},
			mockSetup: func(deps *testDeps) {},
			assertFn: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "blacklist fails",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: generateTestRefreshToken(t)}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					Blacklist(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errs.NewInternal(nil, "redis error"))
			},
			assertFn: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			req := tt.reqFn(t, deps)
			err := deps.svc.Logout(context.Background(), req)
			tt.assertFn(t, err)
		})
	}
}

// --- RefreshToken Tests ---

func TestRefreshToken(t *testing.T) {
	tests := []struct {
		name      string
		reqFn     func(t *testing.T, deps *testDeps) entity.AuthToken
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.AuthToken, err error)
	}{
		{
			name: "success",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: generateTestRefreshToken(t)}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					IsBlacklisted(gomock.Any(), gomock.Any()).
					Return(false, nil)
				deps.tokenCache.EXPECT().
					Blacklist(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				require.NoError(t, err)
				assert.NotNil(t, got)
				assert.NotEmpty(t, got.AccessToken)
				assert.NotEmpty(t, got.RefreshToken)
			},
		},
		{
			name: "blacklisted token",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: generateTestRefreshToken(t)}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					IsBlacklisted(gomock.Any(), gomock.Any()).
					Return(true, nil)
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "invalid token",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: "invalid-token-string"}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					IsBlacklisted(gomock.Any(), gomock.Any()).
					Return(false, nil)
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrUnauthorized, errs.GetCode(err))
			},
		},
		{
			name: "blacklist check fails",
			reqFn: func(t *testing.T, deps *testDeps) entity.AuthToken {
				return entity.AuthToken{RefreshToken: generateTestRefreshToken(t)}
			},
			mockSetup: func(deps *testDeps) {
				deps.tokenCache.EXPECT().
					IsBlacklisted(gomock.Any(), gomock.Any()).
					Return(false, errs.NewInternal(nil, "redis error"))
			},
			assertFn: func(t *testing.T, got *entity.AuthToken, err error) {
				assert.Nil(t, got)
				require.Error(t, err)
				assert.Equal(t, errs.ErrInternal, errs.GetCode(err))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := setupTest(t)
			tt.mockSetup(deps)

			req := tt.reqFn(t, deps)
			got, err := deps.svc.RefreshToken(context.Background(), req)
			tt.assertFn(t, got, err)
		})
	}
}

// --- Helpers ---

// generateTestRefreshToken creates a valid refresh token for testing.
// Must be called after setupTest sets env vars / viper keys.
func generateTestRefreshToken(t *testing.T) string {
	t.Helper()
	token, err := jwt.GenerateToken(jwt.UserCtx{Id: "user1", Email: "user@test.com"}, true)
	require.NoError(t, err)
	return token
}
