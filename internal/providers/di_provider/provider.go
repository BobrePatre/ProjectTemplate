package di_provider

import (
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/grpc/v1/example"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/http/handlers/example/v1"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/config"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/repository"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/service"
)

type DiProvider struct {
	redisConfig *config.RedisConfig

	validate *validator.Validate

	grpcConfig    *config.GRPCConfig
	httpConfig    *config.HttpConfig
	webAuthConfig *config.WebAuthConfig
	appConfig     *config.AppConfig

	exampleRepository repository.ExampleRepository
	exampleService    service.ExampleService
	exampleHandler    *v1.Handler
	exampleImpl       *example.Implementation

	httpAuthMiddlewareConstructor       web_auth_provider.AuthHttpMiddlewareConstructor
	grpcUnaryAuthInterceptorConstructor web_auth_provider.AuthGrpcUnaryInterceptorConstructor
}

func NewDiProvider() *DiProvider {
	return &DiProvider{}
}
