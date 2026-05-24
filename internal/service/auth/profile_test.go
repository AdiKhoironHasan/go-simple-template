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
