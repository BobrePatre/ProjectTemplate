package app

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/providers/di_provider"
)

var _ = (*App)(nil)

func (a *App) initDeps(ctx context.Context) error {
	inits := []configFunc{
		a.initdiProvider,
		a.initLogger,
		a.initRedis,
		a.initWebAuthProvider,
		a.initGRPCServer,
		a.initGatewayServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initdiProvider(_ context.Context) error {
	a.diProvider = di_provider.NewDiProvider()
	return nil
}
