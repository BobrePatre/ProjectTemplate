package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"log"
)

func (p *DiProvider) GrpcConfig() *config.GrpcConfig {
	if p.grpcConfig == nil {
		cfg, err := config.NewGrpcConfig(p.Validate())
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		p.grpcConfig = cfg
	}

	return p.grpcConfig
}
