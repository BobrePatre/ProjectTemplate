package di_provider

import (
	"context"
	grpcMiddlewares "github.com/BobrePatre/ProjectTemplate/internal/api/grpc/interceptors"
	httpMiddlewares "github.com/BobrePatre/ProjectTemplate/internal/api/http/middlewares"
	webAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	keycloakAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/keycloak_redis"
	"log"
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
		p.grpcUnaryAuthInterceptorConstructor = grpcMiddlewares.AuthInerceptor
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
			log.Fatalf("failed to fetch jwks from sso: %s", err.Error())
		}
	}

	return p.webAuthProvider
}
