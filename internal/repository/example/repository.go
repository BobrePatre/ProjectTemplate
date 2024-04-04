package example

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/internal/models"
	def "github.com/BobrePatre/ProjectTemplate/internal/repository"
	"sync"
)

var _ def.ExampleRepository = (*Repository)(nil)

type Repository struct {
	inMemoryStorage []Example
	mutex           sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Example(_ context.Context, example models.Example) (models.Example, error) {
	//Simple in-memory example
	r.mutex.Lock()
	defer r.mutex.Unlock()

	repoModel := RepoFromService(example)
	r.inMemoryStorage = append(r.inMemoryStorage, repoModel)

	return example, nil
}
