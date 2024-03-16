package app

import (
	"context"
	descExample "github.com/BobrePatre/ProjectTemplate/pkg/api/grpc/golang/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

var _ = (*App)(nil)

func (a *App) initGRPCServer(_ context.Context) error {
	authInterceptorConstructor := a.diProvider.GrpcAuthUnaryInterceptorConstructor()

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			authInterceptorConstructor(a.diProvider.WebAuthProvider()),
		),
	)

	reflection.Register(a.grpcServer)

	descExample.RegisterExampleServiceServer(a.grpcServer, a.diProvider.ExampleImpl())

	return nil
}

func (a *App) runGRPCServer() error {
	slog.Info(
		startingMsg,
		grpcServerTag,
		slog.String(addressMsg, a.diProvider.GRPCConfig().Address()),
	)

	listener, err := net.Listen("tcp", a.diProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
