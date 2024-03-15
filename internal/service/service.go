package service

import (
	"context"

	"github.com/BobrePatre/ProjectTemplate/v1/internal/model"
)

type ExampleService interface {
	Example(ctx context.Context, example model.Example) (model.Example, error)
}
