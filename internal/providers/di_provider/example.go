package di_provider

import (
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/grpc/v1/example"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/http/handlers/example/v1"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/repository"
	exampleRepository "github.com/BobrePatre/ProjectTemplate/v1/internal/repository/example"
	"github.com/BobrePatre/ProjectTemplate/v1/internal/service"
	exampleService "github.com/BobrePatre/ProjectTemplate/v1/internal/service/example"
)

var _ = (*DiProvider)(nil)

func (p *DiProvider) ExampleRepository() repository.ExampleRepository {
	if p.exampleRepository == nil {
		p.exampleRepository = exampleRepository.NewRepository()
	}
	return p.exampleRepository
}

func (p *DiProvider) ExampleService() service.ExampleService {
	if p.exampleService == nil {
		p.exampleService = exampleService.NewService(p.ExampleRepository())
	}
	return p.exampleService
}

func (p *DiProvider) ExampleHandler() *v1.Handler {
	if p.exampleHandler == nil {
		p.exampleHandler = v1.NewHandler(p.ExampleService())
	}
	return p.exampleHandler
}

func (p *DiProvider) ExampleImpl() *example.Implementation {
	if p.exampleImpl == nil {
		p.exampleImpl = example.NewImplementation(p.ExampleService())
	}
	return p.exampleImpl
}
