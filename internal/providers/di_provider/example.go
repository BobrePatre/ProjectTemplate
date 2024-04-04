package diProvider

import (
	exampleGrpc "github.com/BobrePatre/ProjectTemplate/internal/delivery/grpc/impl/example"
	exampleHttp "github.com/BobrePatre/ProjectTemplate/internal/delivery/http/handlers/example"
	"github.com/BobrePatre/ProjectTemplate/internal/repository"
	exampleRepository "github.com/BobrePatre/ProjectTemplate/internal/repository/example"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	exampleService "github.com/BobrePatre/ProjectTemplate/internal/service/example"
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

func (p *DiProvider) ExampleHandler() *exampleHttp.Handler {
	if p.exampleHandler == nil {
		p.exampleHandler = exampleHttp.NewHandler(p.ExampleService())
	}
	return p.exampleHandler
}

func (p *DiProvider) ExampleImpl() *exampleGrpc.Implementation {
	if p.exampleImpl == nil {
		p.exampleImpl = exampleGrpc.NewImplementation(p.ExampleService())
	}
	return p.exampleImpl
}
