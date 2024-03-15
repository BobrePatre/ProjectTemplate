package web_auth_provider

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"google.golang.org/grpc"
	"time"
)

type WebAuthProvider interface {
	VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, error)
	TokenKeyfunc(ctx context.Context) jwt.Keyfunc
	FetchJwkSet(ctx context.Context) (jwk.Set, error)
	IsUserHaveRoles(roles []string, userRoles []string) bool
	SerializeJwkSet(key jwk.Set) (string, error)
	DeserializeJwkSet(serializedKey string) (jwk.Set, error)
	Authorize(ctx context.Context, tokenString string, roles []string) (model.UserDetails, error)
	CheckSsoConnection(ctx context.Context) error
}

const (
	JwkKeySet      = "jwk-set"
	UserDetailsKey = "userDetails"
)

type JwkOptions struct {
	RefreshJwkTimeout time.Duration
	JwkPublicUri      string
}

type AuthHttpMiddleware func(roles ...string) gin.HandlerFunc
type AuthHttpMiddlewareConstructor func(provider WebAuthProvider) func(roles ...string) gin.HandlerFunc
type AuthGrpcUnaryInterceptorConstructor func(provider WebAuthProvider) grpc.UnaryServerInterceptor
