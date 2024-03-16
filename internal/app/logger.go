package app

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/constants"
	"log/slog"
	"os"
)

var _ = (*App)(nil)

func (a *App) initLogger(_ context.Context) error {
	var logHandler slog.Handler

	switch a.diProvider.AppConfig().ENV {
	case constants.EnvDevelopment:
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	case constants.EnvProduction:
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		})
	}

	a.logger = slog.New(logHandler)
	slog.SetDefault(a.logger)
	return nil
}
