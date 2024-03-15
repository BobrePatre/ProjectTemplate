package example

import "github.com/BobrePatre/ProjectTemplate/v1/internal/model"

func RepoFromService(example model.Example) Example {
	return Example{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
