package keycloak_redis

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/lestrrat-go/jwx/jwk"
	"log/slog"
)

var _ webAuthProvider.WebAuthProvider = (*Provider)(nil)

func (p *Provider) FetchJwkSet(ctx context.Context) (jwk.Set, error) {

	result, err := p.redis.Get(ctx, webAuthProvider.JwkKeySet).Result()
	if err == nil {
		slog.Info("Jwk get from cache")
		resultSet, err := p.DeserializeJwkSet(result)
		if err != nil {
			return nil, err
		}
		return resultSet, nil
	}

	resultSet, err := jwk.Fetch(ctx, p.jwkOpts.JwkPublicUri)
	if err != nil {
		return nil, err
	}

	slog.Info("Jwk get from remote")
	serializedKeySet, err := p.SerializeJwkSet(resultSet)
	if err != nil {
		return nil, err
	}

	err = p.redis.Set(ctx, webAuthProvider.JwkKeySet, serializedKeySet, p.jwkOpts.RefreshJwkTimeout).Err()
	if err != nil {
		return nil, err
	}

	return resultSet, nil

}
