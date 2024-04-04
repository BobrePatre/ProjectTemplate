package example

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/models"
	"github.com/BobrePatre/ProjectTemplate/internal/repository"
	def "github.com/BobrePatre/ProjectTemplate/internal/service"
)

var _ def.ExampleService = (*Service)(nil)

type Service struct {
	repository repository.ExampleRepository
}

func NewService(repository repository.ExampleRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Example(ctx context.Context, example models.Example) (models.Example, error) {

	if len(example.Title) < 4 {
		return models.Example{}, models.ErrorExampleTitleTooShort
	}

	m, err := s.repository.Example(ctx, example)
	if err != nil {
		return models.Example{}, err
	}
	return m, nil
}
