package context

import (
	"context"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/consts"
)

func SetUserCtx(ctx context.Context, userCtx *entity.UserCtx) context.Context {
	return context.WithValue(ctx, consts.CtxUser, userCtx)
}

// GetUserCtx extracts the user context from the request context
func GetUserCtx(ctx context.Context) (*entity.UserCtx, bool) {
	userCtx, ok := ctx.Value(consts.CtxUser).(*entity.UserCtx)
	return userCtx, ok
}
