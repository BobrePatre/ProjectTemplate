package keycloak_redis

import (
	"github.com/goccy/go-json"
	"github.com/lestrrat-go/jwx/jwk"
	"log/slog"
)

//func (p *Provider) SerializePublicKey(key rsa.PublicKey) (string, error) {
//	serialized := map[string]string{
//		"N": key.N.String(),
//		"E": fmt.Sprintf("%d", key.E),
//	}
//	serializedKey, err := json.Marshal(serialized)
//	if err != nil {
//		return "", err
//	}
//	return string(serializedKey), nil
//}
//
//func (p *Provider) DeserializePublicKey(serializedKey string) (rsa.PublicKey, error) {
//	var serialized map[string]string
//	if err := json.Unmarshal([]byte(serializedKey), &serialized); err != nil {
//		return rsa.PublicKey{}, err
//	}
//
//	N := new(big.Int)
//	N.SetString(serialized["N"], 10)
//
//	E, err := strconv.ParseInt(serialized["E"], 10, 64)
//	if err != nil {
//		return rsa.PublicKey{}, err
//	}
//
//	return rsa.PublicKey{
//		N: N,
//		E: int(E),
//	}, nil
//}

// SerializeJwkSet сериализует jwk.Set в строку JSON.
func (p *Provider) SerializeJwkSet(key jwk.Set) (string, error) {
	// Сериализуем ключи JWK в строку JSON.
	serializedKey, err := json.Marshal(key)
	if err != nil {
		slog.Error("Failed to serialize JWK set", slog.String("err", err.Error()))
		return "", err
	}
	return string(serializedKey), nil
}

// DeserializeJwkSet десериализует строку JSON в jwk.Set.
func (p *Provider) DeserializeJwkSet(serializedKey string) (jwk.Set, error) {
	keySet, err := jwk.Parse([]byte(serializedKey))
	if err != nil {
		slog.Error("Failed to deserialize JWK set", slog.String("err", err.Error()))
		return nil, err
	}
	return keySet, nil
}
