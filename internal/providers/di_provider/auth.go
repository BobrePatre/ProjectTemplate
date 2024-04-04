package diProvider

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/grpc/interceptors"
	httpMiddlewares "github.com/BobrePatre/ProjectTemplate/internal/delivery/http/middlewares"
	webAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	keycloakAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/keycloak_redis"
	"log"
	"log/slog"
	"os"
	"time"
)

func (p *DiProvider) HttpAuthMiddlewareConstructor() webAuthProvider.AuthHttpMiddlewareConstructor {
	if p.httpAuthMiddlewareConstructor == nil {
		p.httpAuthMiddlewareConstructor = httpMiddlewares.AuthMiddleware
	}
	return p.httpAuthMiddlewareConstructor

}

func (p *DiProvider) GrpcAuthUnaryInterceptorConstructor() webAuthProvider.AuthGrpcUnaryInterceptorConstructor {
	if p.grpcUnaryAuthInterceptorConstructor == nil {
		p.grpcUnaryAuthInterceptorConstructor = interceptors.AuthInerceptor
	}
	return p.grpcUnaryAuthInterceptorConstructor
}

func (p *DiProvider) WebAuthProvider() webAuthProvider.WebAuthProvider {
	if p.webAuthProvider == nil {
		p.webAuthProvider = keycloakAuthProvider.NewProvider(p.RedisClient(), webAuthProvider.JwkOptions{
			JwkPublicUri:      p.AuthConfig().PublicJwkUri,
			RefreshJwkTimeout: p.AuthConfig().RefreshJwkTimeout,
		}, p.Validate(), p.AuthConfig().ClientId)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := p.webAuthProvider.CheckSsoConnection(ctx)
		if err != nil {
			slog.Error("failed to fetch jwks from sso", "error", err.Error())
			os.Exit(1)
		}
	}

	return p.webAuthProvider
}

func (p *DiProvider) AuthConfig() config.WebAuthConfig {
	if p.webAuthConfig == nil {
		cfg, err := config.NewAuthConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}
		p.webAuthConfig = cfg
	}
	return *p.webAuthConfig
}
