package example

import "github.com/BobrePatre/ProjectTemplate/internal/model"

func RepoFromService(example model.Example) Example {
	return Example{
		Title:       example.Title,
		Description: example.Description,
		Body:        example.Body,
	}
}
