package app

import (
	"context"
	descExample "github.com/BobrePatre/ProjectTemplate/pkg/api/golang/example"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ = (*App)(nil)

func (a *App) initGatewayServer(ctx context.Context) error {
	a.gatewayServer = runtime.NewServeMux()

	err := descExample.RegisterExampleServiceHandlerFromEndpoint(ctx, a.gatewayServer, a.diProvider.GRPCConfig().Address(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return err
	}
	return nil
}
