package repository

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/model"
)

type ExampleRepository interface {
	Example(context.Context, model.Example) (model.Example, error)
}
