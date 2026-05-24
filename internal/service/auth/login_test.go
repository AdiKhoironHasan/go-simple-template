package auth

import (
	"context"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/crypto"
	"go-simple-template/internal/pkg/errs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

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
