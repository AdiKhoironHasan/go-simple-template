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
