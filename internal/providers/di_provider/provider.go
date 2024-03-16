package di_provider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/api/grpc/example"
	example2 "github.com/BobrePatre/ProjectTemplate/internal/api/http/handlers/example"
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	webAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/BobrePatre/ProjectTemplate/internal/repository"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type DiProvider struct {
	redisClient *redis.Client
	redisConfig *config.RedisConfig

	validate *validator.Validate

	grpcConfig *config.GRPCConfig
	httpConfig *config.HttpConfig
	appConfig  *config.AppConfig

	exampleRepository repository.ExampleRepository
	exampleService    service.ExampleService
	exampleHandler    *example2.Handler
	exampleImpl       *example.Implementation

	webAuthProvider webAuthProvider.WebAuthProvider
	webAuthConfig   *config.WebAuthConfig

	httpAuthMiddlewareConstructor       webAuthProvider.AuthHttpMiddlewareConstructor
	grpcUnaryAuthInterceptorConstructor webAuthProvider.AuthGrpcUnaryInterceptorConstructor
}

func NewDiProvider() *DiProvider {
	return &DiProvider{}
}
