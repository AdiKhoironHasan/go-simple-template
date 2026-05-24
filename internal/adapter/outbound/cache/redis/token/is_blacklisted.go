package token

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"

	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"
)

func (c *cacheRepo) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	key := blacklistPrefix + hashToken(token)
	exists, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to check token blacklist", slog.String(consts.Error, err.Error()))
		return false, err
	}
	return exists > 0, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
