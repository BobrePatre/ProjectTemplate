package di_provider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"log"
)

var _ = (*DiProvider)(nil)

func (p *DiProvider) GRPCConfig() config.GRPCConfig {
	if p.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		p.grpcConfig = cfg
	}

	return *p.grpcConfig
}

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

func (p *DiProvider) HTTPConfig() config.HttpConfig {
	if p.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		p.httpConfig = cfg
	}

	return *p.httpConfig
}

func (p *DiProvider) AppConfig() config.AppConfig {
	if p.appConfig == nil {
		cfg, err := config.NewAppConfig()
		if err != nil {
			log.Fatalf("failed to get app config: %s", err.Error())
		}
		p.appConfig = cfg
	}

	return *p.appConfig
}

func (p *DiProvider) AuthConfig() config.WebAuthConfig {
	if p.webAuthConfig == nil {
		cfg, err := config.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}
		p.webAuthConfig = cfg
	}
	return *p.webAuthConfig
}
