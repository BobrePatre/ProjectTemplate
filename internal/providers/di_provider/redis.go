package diProvider

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
	"time"
)

const (
	redisSuccessResponse = "PONG"
)

func (p *DiProvider) RedisConfig() config.RedisConfig {
	if p.redisConfig == nil {
		redisConfig, err := config.NewRedisConfig(p.Validate())
		if err != nil {
			slog.Error("failed to get redis config", "err", err.Error())
			os.Exit(1)
		}
		p.redisConfig = redisConfig
	}
	return *p.redisConfig
}

func (p *DiProvider) RedisClient() *redis.Client {
	if p.redisClient == nil {
		p.redisClient = redis.NewClient(&redis.Options{
			Addr:     p.RedisConfig().Address(),
			Password: p.RedisConfig().Password,
			DB:       p.RedisConfig().Database,
		})

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		status := p.redisClient.Ping(timeoutCtx)
		result, err := status.Result()
		if err != nil {
			slog.Error("failed to connect to redis", "error", err.Error())
			os.Exit(1)
		}

		if result != redisSuccessResponse {
			slog.Error("failed to connect to redis", "error", err.Error())
			os.Exit(1)
		}
	}

	return p.redisClient

}
