package auth

import (
	"context"
	"testing"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

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
