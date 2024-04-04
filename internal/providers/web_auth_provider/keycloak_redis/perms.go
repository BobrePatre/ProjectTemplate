package keycloak_redis

import "github.com/BobrePatre/ProjectTemplate/internal/utils/helpers"

func (p *Provider) IsUserHaveRoles(roles []string, userRoles []string) bool {
	if len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if helpers.IsArrayContains(userRoles, role) {
			return true
		}
	}
	return false
}
