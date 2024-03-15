package keycloak_redis

import "github.com/BobrePatre/ProjectTemplate/v1/internal/utils/helpers"

func (p *Provider) IsUserHaveRoles(roles []string, userRoles []string) bool {
	for _, role := range roles {
		if helpers.IsArrayContains(userRoles, role) {
			return true
		}
	}
	return false
}
