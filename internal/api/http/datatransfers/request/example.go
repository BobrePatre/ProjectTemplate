package request

type ExampleRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Body        string `json:"body" validate:"required"`
}
