package converters

import (
	"github.com/BobrePatre/ProjectTemplate/internal/model"
	desc "github.com/BobrePatre/ProjectTemplate/pkg/api/golang/example/v2"
)

func ExampleFromProtoRequest(request *desc.ExampleRequest) model.Example {
	return model.Example{
		Title:       request.Title,
		Description: request.Description,
		Body:        request.Body,
	}
}

func ExampleResponseFromModel(example model.Example) *desc.ExampleResponse {
	return &desc.ExampleResponse{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}