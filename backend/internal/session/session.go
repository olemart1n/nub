package session

import (
	"context"
	"fmt"
	"time"
)

func SetSession(ctx context.Context, sessionID string, userID string, ttl time.Duration) error {
	return Client.Set(ctx, fmt.Sprintf("session:%s", sessionID), userID, ttl).Err()
}

func GetSession(ctx context.Context, sessionID string) (string, error) {
	return Client.Get(ctx, fmt.Sprintf("session:%s", sessionID)).Result()
}

func DeleteSession(ctx context.Context, sessionID string) error {
	return Client.Del(ctx, fmt.Sprintf("session:%s", sessionID)).Err()
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value("userID").(string)
	return userID, ok
}
