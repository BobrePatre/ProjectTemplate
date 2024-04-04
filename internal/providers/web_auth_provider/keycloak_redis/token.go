package keycloak_redis

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	authErrors "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/models"
	"github.com/golang-jwt/jwt"
	"log/slog"
)

var _ webAuthProvider.WebAuthProvider = (*Provider)(nil)

func (p *Provider) VerifyToken(ctx context.Context, tokenString string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(tokenString, p.TokenKeyfunc(ctx))

	if err != nil {
		slog.Error("Failed to parse token", slog.String("err", err.Error()))
		return token, authErrors.InvalidTokenError
	}

	return token, nil
}

func (p *Provider) TokenKeyfunc(ctx context.Context) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		var rawKey rsa.PublicKey

		keySet, err := p.FetchJwkSet(ctx)
		if err != nil {
			slog.Error("Failed to get jwk set", slog.String("err", err.Error()))
			return nil, err
		}

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf(authErrors.ValidationTokenError.Error(), err)
		}

		keyID, ok := token.Header["kid"].(string)
		if !ok {
			err = fmt.Errorf("expecting JWT header to have string 'kid'")
			return nil, fmt.Errorf(authErrors.ValidationTokenError.Error(), err)
		}

		key, found := keySet.LookupKeyID(keyID)
		if !found {
			err = fmt.Errorf("unable to find key")
			return nil, fmt.Errorf(authErrors.JwkKetNotFound.Error(), err)
		}

		err = key.Raw(&rawKey)
		if err != nil {
			slog.Error("Failed to get raw key", slog.String("err", err.Error()))
			return rawKey, authErrors.InvalidTokenError
		}

		return &rawKey, err
	}
}
