package repository

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/model"
)

type ExampleRepository interface {
	Example(context.Context, model.Example) (model.Example, error)
}
