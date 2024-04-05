package app

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/constants"
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/http/middlewares"
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/http/routes/example"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

var _ = (*App)(nil)

func (a *App) initHTTPServer(_ context.Context) error {

	switch a.diProvider.AppConfig().ENV {
	case constants.EnvDevelopment:
		gin.SetMode(gin.DebugMode)
	case constants.EnvProduction:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.SlogLoggerMiddleware())

	// TODO: create config for cors
	router.Use(cors.Default())

	authMiddlewareConstructor := a.diProvider.HttpAuthMiddlewareConstructor()

	v1RouterGroup := router.Group("/api/v1")
	example.NewRouter(v1RouterGroup,
		authMiddlewareConstructor(a.diProvider.WebAuthProvider()),
		a.diProvider.ExampleHandler(),
	).RegisterRoutes()

	router.Any("/grpc/v1/*any", gin.WrapH(a.gatewayServer))

	a.httpServer = &http.Server{
		Addr:    a.diProvider.HTTPConfig().Address(),
		Handler: router,
	}
	return nil
}

func (a *App) runHTTPServer() error {
	slog.Info(
		startingMsg,
		httpServerTag,
		slog.String(addressMsg, a.diProvider.HTTPConfig().Address()),
	)
	return a.httpServer.ListenAndServe()
}
