package di_provider

import (
	grpcMiddlewares "github.com/BobrePatre/ProjectTemplate/v1/internal/api/grpc/interceptors"
	httpMiddlewares "github.com/BobrePatre/ProjectTemplate/v1/internal/api/http/middlewares"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
)

func (p *DiProvider) HttpAuthMiddlewareConstructor() web_auth_provider.AuthHttpMiddlewareConstructor {
	if p.httpAuthMiddlewareConstructor == nil {
		p.httpAuthMiddlewareConstructor = httpMiddlewares.AuthMiddleware
	}
	return p.httpAuthMiddlewareConstructor

}

func (p *DiProvider) GrpcAuthUnaryInterceptorConstructor() web_auth_provider.AuthGrpcUnaryInterceptorConstructor {
	if p.grpcUnaryAuthInterceptorConstructor == nil {
		p.grpcUnaryAuthInterceptorConstructor = grpcMiddlewares.AuthInerceptor
	}
	return p.grpcUnaryAuthInterceptorConstructor
}
