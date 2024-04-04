package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"log"
)

func (p *DiProvider) AppConfig() config.AppConfig {
	if p.appConfig == nil {
		cfg, err := config.NewAppConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get app config: %s", err.Error())
		}
		p.appConfig = cfg
	}

	return *p.appConfig
}
