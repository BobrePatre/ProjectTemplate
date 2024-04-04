package keycloak_redis

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/lestrrat-go/jwx/jwk"
)

var _ webAuthProvider.WebAuthProvider = (*Provider)(nil)

func (p *Provider) CheckSsoConnection(ctx context.Context) error {

	_, err := jwk.Fetch(ctx, p.jwkOpts.JwkPublicUri)
	if err != nil {
		return err
	}

	return nil
}
