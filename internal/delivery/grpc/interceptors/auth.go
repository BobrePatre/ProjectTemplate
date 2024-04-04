package interceptors

import (
	"context"
	"errors"
	authProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	authErrors "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
	"strings"
)

func AuthInerceptor(provider authProvider.WebAuthProvider) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		slog.Debug("auth interceptor", slog.String("method", info.FullMethod))
		switch info.FullMethod {
		case "/v1.example.ExampleService/Example":
			token, err := parseTokenFromHeader(ctx)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			userDetails, err := provider.Authorize(ctx, token, []string{"user"})
			if err != nil {
				switch {
				case errors.Is(err, authErrors.AccessDeniedError):
					return nil, status.Error(codes.PermissionDenied, err.Error())
				default:
					return nil, status.Error(codes.Unauthenticated, err.Error())
				}
			}

			ctx = context.WithValue(ctx, authProvider.UserDetailsKey, userDetails)
			return handler(ctx, req)
		default:
			return nil, status.Error(codes.Unavailable, "method dont have security config")
		}

	}
}

func parseTokenFromHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadata is not provided")
	}
	authHeader, ok := md["authorization"]
	if !ok {
		return "", errors.New("authorization header is not provided")
	}
	if len(authHeader) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	headerArr := strings.Split(authHeader[0], " ")
	if len(headerArr) != 2 || headerArr[0] != "Bearer" {
		slog.Error("Invalid token format")
		return "", errors.New("invalid token format, use Bearer {token}")
	}
	return headerArr[1], nil
}
