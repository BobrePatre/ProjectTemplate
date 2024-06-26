package app

import (
	"context"
	"errors"
	diProvider "github.com/BobrePatre/ProjectTemplate/internal/providers/di_provider"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type configFunc func(context.Context) error

type App struct {
	diProvider    *diProvider.DiProvider
	grpcServer    *grpc.Server
	gatewayServer *runtime.ServeMux
	httpServer    *http.Server
	logger        *slog.Logger
}

const (
	startingMsg = "starting"
	stopMsg     = "stopping"
	addressMsg  = "address"
)

var (
	httpServerTag = slog.String("server", "http")
	grpcServerTag = slog.String("server", "grpc")
)

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		if err := a.runGRPCServer(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runHTTPServer(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info(stopMsg, grpcServerTag)
	a.grpcServer.GracefulStop()

	slog.Info(stopMsg, httpServerTag)
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	wg.Wait()

	slog.Info("Servers shutdown successfully")
}
