package example

import (
	"github.com/BobrePatre/ProjectTemplate/internal/api/http/converters"
	"github.com/BobrePatre/ProjectTemplate/internal/api/http/datatransfers/request"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service service.ExampleService
}

func NewHandler(
	service service.ExampleService,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Example(ctx *gin.Context) {

	var req request.ExampleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := validator.New().Struct(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	result, err := h.service.Example(ctx, converters.ToExampleFromRequest(req))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, converters.ToExampleFromService(result))
}
