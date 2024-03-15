package example

import (
	"context"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/model"
	def "github.com/BobrePatre/ProjectTemplate/v1/internal/repository"
	"sync"
)

var _ def.ExampleRepository = (*Repository)(nil)

type Repository struct {
	inMemoryStorage []Example
	mutex           *sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		mutex: &sync.RWMutex{},
	}
}

func (r Repository) Example(_ context.Context, example model.Example) (model.Example, error) {
	//Simple in-memory example
	r.mutex.Lock()
	defer r.mutex.Unlock()

	repoModel := RepoFromService(example)
	r.inMemoryStorage = append(r.inMemoryStorage, repoModel)

	return example, nil
}
