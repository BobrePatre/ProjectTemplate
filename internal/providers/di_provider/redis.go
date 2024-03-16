package di_provider

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func (p *DiProvider) RedisConfig() config.RedisConfig {
	if p.redisConfig == nil {
		redisConfig, err := config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
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
			log.Fatalf("failed to connect to redis: %s", err.Error())
		}

		if result != "PONG" {
			log.Fatalf("failed to connect to redis: %s", err.Error())
		}
	}

	return p.redisClient

}
