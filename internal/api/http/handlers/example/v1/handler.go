package v1

import (
	"github.com/BobrePatre/ProjectTemplate/internal/api/http/converters"
	"github.com/BobrePatre/ProjectTemplate/internal/api/http/datatransfers/request/v1"
	"github.com/BobrePatre/ProjectTemplate/internal/service"
	"github.com/gin-gonic/gin"
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

	var req v1.ExampleRequest
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
