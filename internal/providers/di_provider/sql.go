package diProvider

import (
	"github.com/BobrePatre/ProjectTemplate/internal/config"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
)

func (p *DiProvider) PostgresqlConfig() *config.PostgresqlConfig {
	if p.postgresqlConfig == nil {
		cfg, err := config.NewPostgresConfig(p.Validate())
		if err != nil {
			slog.Error("failed to get postgresql config", "err", err.Error())
			os.Exit(1)
		}
		p.postgresqlConfig = cfg
	}

	return p.postgresqlConfig
}

func (p *DiProvider) SqlDatabase() *sqlx.DB {
	if p.sqlDatabase == nil {
		db, err := sqlx.Connect("postgres", p.PostgresqlConfig().Datasource("disable"))
		slog.Debug("postgres config", "config", p.PostgresqlConfig())
		if err != nil {
			slog.Error("failed to connect to sql database", "err", err.Error())
			os.Exit(1)
		}
		p.sqlDatabase = db
	}

	return p.sqlDatabase
}
