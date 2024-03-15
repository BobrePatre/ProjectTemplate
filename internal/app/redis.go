package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ = (*App)(nil)

func (a *App) initRedis(ctx context.Context) error {
	a.redis = redis.NewClient(&redis.Options{
		Addr:     a.diProvider.RedisConfig().Address(),
		Password: a.diProvider.RedisConfig().Password,
		DB:       a.diProvider.RedisConfig().Database,
	})
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	status := a.redis.Ping(timeoutCtx)
	result, err := status.Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %s", err.Error())
	}

	if result != "PONG" {
		return fmt.Errorf("failed to connect to redis: %s", err.Error())
	}

	return nil
}
