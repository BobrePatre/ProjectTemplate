package repository

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/models"
)

type (
	ExampleRepository interface {
		Example(context.Context, models.Example) (models.Example, error)
	}
)
