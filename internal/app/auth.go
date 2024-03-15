package app

import (
	"context"
	"fmt"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	keycloakAuthProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/keycloak_redis"
	"time"
)

func (a *App) initWebAuthProvider(_ context.Context) error {
	a.authProvider = keycloakAuthProvider.NewProvider(a.redis, web_auth_provider.JwkOptions{
		JwkPublicUri:      a.diProvider.AuthConfig().PublicJwkUri,
		RefreshJwkTimeout: a.diProvider.AuthConfig().RefreshJwkTimeout,
	}, a.diProvider.Validate(), a.diProvider.AuthConfig().ClientId)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := a.authProvider.CheckSsoConnection(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch jwks from sso: %s", err.Error())
	}

	return nil
}
