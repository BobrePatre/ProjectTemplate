package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"log"
)

func (p *DiProvider) HTTPConfig() *config.HttpConfig {
	if p.httpConfig == nil {
		cfg, err := config.NewHTTPConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		p.httpConfig = cfg
	}

	return p.httpConfig
}
