package auth

import (
	"context"
	"testing"

	"github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRefreshToken(t *testing.T) {
	tests := []struct {
		name      string
		req       entity.AuthToken
		mockSetup func(deps *testDeps)
		assertFn  func(t *testing.T, got *entity.AuthToken, err error)
	}{
		{
			name: "success",
			req: entity.AuthToken{
				RefreshToken: generateTestRefreshToken(t),
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
			req: entity.AuthToken{
				RefreshToken: generateTestRefreshToken(t),
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
			req: entity.AuthToken{
				RefreshToken: "invalid-token-string",
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
			req: entity.AuthToken{
				RefreshToken: generateTestRefreshToken(t),
				// RefreshToken: "any-token", --- IGNORE ---
				// We want to test the error path of blacklist check, so token value doesn't matter as long as it's not blacklisted.
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

			got, err := deps.svc.RefreshToken(context.Background(), tt.req)
			tt.assertFn(t, got, err)
		})
	}
}
