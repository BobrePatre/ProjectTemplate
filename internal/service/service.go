package service

import (
	"context"

	"github.com/BobrePatre/ProjectTemplate/internal/model"
)

type ExampleService interface {
	Example(ctx context.Context, example model.Example) (model.Example, error)
}
