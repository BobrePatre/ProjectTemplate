package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	grpcExample "github.com/BobrePatre/ProjectTemplate/internal/delivery/grpc/impl/example"
	httpExample "github.com/BobrePatre/ProjectTemplate/internal/delivery/http/handlers/example"
	webAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/BobrePatre/ProjectTemplate/internal/repository"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DiProvider struct {
	redisClient *redis.Client
	redisConfig *config.RedisConfig

	sqlDatabase      *sqlx.DB
	postgresqlConfig *config.PostgresqlConfig

	validate *validator.Validate

	httpConfig *config.HttpConfig
	grpcConfig *config.GrpcConfig
	appConfig  *config.AppConfig

	webAuthProvider webAuthProvider.WebAuthProvider
	webAuthConfig   *config.WebAuthConfig

	httpAuthMiddlewareConstructor       webAuthProvider.AuthHttpMiddlewareConstructor
	grpcUnaryAuthInterceptorConstructor webAuthProvider.AuthGrpcUnaryInterceptorConstructor

	exampleRepository repository.ExampleRepository
	exampleService    service.ExampleService
	exampleHandler    *httpExample.Handler
	exampleImpl       *grpcExample.Implementation
}

func NewDiProvider() *DiProvider {
	return &DiProvider{}
}
