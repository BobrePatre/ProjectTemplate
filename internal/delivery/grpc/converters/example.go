package converters

import (
	"github.com/BobrePatre/ProjectTemplate/internal/models"
	desc "github.com/BobrePatre/ProjectTemplate/pkg/api/grpc/golang/v1/example"
)

func ExampleFromProtoRequest(request *desc.ExampleRequest) models.Example {
	return models.Example{
		Title:       request.Title,
		Description: request.Description,
		Body:        request.Body,
	}
}

func ExampleResponseFromModel(example models.Example) *desc.ExampleResponse {
	return &desc.ExampleResponse{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
