package service

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/models"
)

type (
	ExampleService interface {
		Example(ctx context.Context, example models.Example) (models.Example, error)
	}
)
