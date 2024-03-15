package di_provider

import "github.com/go-playground/validator/v10"

func (p *DiProvider) Validate() *validator.Validate {
	if p.validate == nil {
		p.validate = validator.New()
	}
	return p.validate
}
