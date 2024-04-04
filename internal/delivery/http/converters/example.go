package converters

import (
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/http/datatransfers/requests"
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/http/datatransfers/responses"
	"github.com/BobrePatre/ProjectTemplate/internal/models"
)

func ToExampleFromRequest(req requests.ExampleRequest) models.Example {
	return models.Example{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	}
}

func ToExampleFromService(example models.Example) responses.ExampleResponse {
	return responses.ExampleResponse{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
