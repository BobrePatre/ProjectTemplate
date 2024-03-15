package keycloak_redis

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
	"github.com/lestrrat-go/jwx/jwk"
)

var _ web_auth_provider.WebAuthProvider = (*Provider)(nil)

func (p *Provider) CheckSsoConnection(ctx context.Context) error {

	_, err := jwk.Fetch(ctx, p.jwkOpts.JwkPublicUri)
	if err != nil {
		return err
	}

	return nil
}
