package example

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/grpc/converters"
	authProvider "github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
	authModels "github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider/model"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/service"
	desc "github.com/BobrePatre/ProjectTemplate/v1/pkg/go/api/example/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Implementation struct {
	desc.UnimplementedExampleServiceServer
	service service.ExampleService
}

func NewImplementation(exampleService service.ExampleService) *Implementation {
	return &Implementation{
		service: exampleService,
	}
}

func (i *Implementation) Example(ctx context.Context, req *desc.ExampleRequest) (*desc.ExampleResponse, error) {

	userDetails := ctx.Value(authProvider.UserDetailsKey).(authModels.UserDetails)
	slog.Debug("user data", slog.Any("user", userDetails))
	result, err := i.service.Example(ctx, converters.ExampleFromProtoRequest(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return converters.ExampleResponseFromModel(result), nil
}
