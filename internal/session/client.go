package session

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis(addr string) error {
	Client = redis.NewClient(&redis.Options{
		Addr: addr, // e.g. "localhost:6379"
	})

	// Ping to verify connection
	if err := Client.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Println("✅ Connected to Redis")
	return nil
}
