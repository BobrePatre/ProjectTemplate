package converters

import (
	requests "github.com/BobrePatre/ProjectTemplate/internal/api/http/datatransfers/request/v1"
	"github.com/BobrePatre/ProjectTemplate/internal/api/http/datatransfers/response"
	"github.com/BobrePatre/ProjectTemplate/internal/model"
)

func ToExampleFromRequest(req requests.ExampleRequest) model.Example {
	return model.Example{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	}
}

func ToExampleFromService(example model.Example) response.ExampleResponse {
	return response.ExampleResponse{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
