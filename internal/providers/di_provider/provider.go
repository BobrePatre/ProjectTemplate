package di_provider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/api/grpc/example"
	example2 "github.com/BobrePatre/ProjectTemplate/internal/api/http/handlers/example"
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/BobrePatre/ProjectTemplate/internal/repository"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	"github.com/go-playground/validator/v10"
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
	exampleHandler    *example2.Handler
	exampleImpl       *example.Implementation

	httpAuthMiddlewareConstructor       web_auth_provider.AuthHttpMiddlewareConstructor
	grpcUnaryAuthInterceptorConstructor web_auth_provider.AuthGrpcUnaryInterceptorConstructor
}

func NewDiProvider() *DiProvider {
	return &DiProvider{}
}
