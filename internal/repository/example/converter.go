package example

import (
	"github.com/BobrePatre/ProjectTemplate/internal/models"
)

func RepoFromService(example models.Example) Example {
	return Example{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
