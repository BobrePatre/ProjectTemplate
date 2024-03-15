package keycloak_redis

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider/model"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"log/slog"
)

var _ web_auth_provider.WebAuthProvider = (*Provider)(nil)

type Provider struct {
	redis    *redis.Client
	jwkOpts  web_auth_provider.JwkOptions
	validate *validator.Validate
	clientID string
}

func NewProvider(redis *redis.Client, jwkOpts web_auth_provider.JwkOptions, validate *validator.Validate, clientID string) *Provider {
	return &Provider{
		redis:    redis,
		jwkOpts:  jwkOpts,
		validate: validate,
		clientID: clientID,
	}
}

func (p *Provider) Authorize(ctx context.Context, tokenString string, neededRoles []string) (model.UserDetails, error) {
	token, err := p.VerifyToken(ctx, tokenString)
	if err != nil {
		slog.Error("failed to verify token", slog.String("err", err.Error()))
		return model.UserDetails{}, model.InvalidTokenError
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {
		slog.Error("failed to get claims", slog.String("err", err.Error()))
		return model.UserDetails{}, model.InvalidTokenError
	}

	if claims["sub"] == "" || claims["sub"] == nil {
		slog.Error("failed to validate sub claim", slog.String("err", err.Error()))
		return model.UserDetails{}, model.InvalidTokenError
	}

	err = p.validate.Var(claims["sub"], "uuid4")
	if err != nil {
		slog.Error("failed to validate sub claim", slog.String("err", err.Error()))
		return model.UserDetails{}, err
	}

	var userRoles []string
	if resourceAccess, ok := claims["resource_access"].(map[string]interface{}); ok {
		if authClient, ok := resourceAccess[p.clientID].(map[string]interface{}); ok {
			if err := mapstructure.Decode(authClient["roles"], &userRoles); err != nil {
				slog.Error("cannot get user roles", slog.String("err", err.Error()))
				userRoles = []string{}
			}
		}
	}

	userEmail, ok := claims["email"].(string)
	if !ok {
		userEmail = ""
	}

	userDetails := model.UserDetails{
		Roles:      userRoles,
		UserId:     claims["sub"].(string),
		Email:      userEmail,
		Username:   claims["preferred_username"].(string),
		Name:       claims["name"].(string),
		FamilyName: claims["family_name"].(string),
	}

	if !p.IsUserHaveRoles(neededRoles, userRoles) {
		slog.Error("user data", slog.Any("userDetails", userDetails))
		slog.Error("user doesn't have needed roles", slog.Any("neededRoles", neededRoles), slog.Any("userRoles", userRoles))
		return userDetails, model.AccessDeniedError
	}

	return userDetails, nil
}
