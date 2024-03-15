package v1

import (
	"github.com/BobrePatre/ProjectTemplate/v1/internal/api/http/handlers/example/v1"
	authProvider "github.com/BobrePatre/ProjectTemplate/v1/internal/providers/web_auth_provider"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router         *gin.RouterGroup
	handler        *v1.Handler
	authMiddleware authProvider.AuthHttpMiddleware
}

func NewRouter(router *gin.RouterGroup, authMiddleware authProvider.AuthHttpMiddleware, handler *v1.Handler) *Router {

	return &Router{
		router:         router,
		handler:        handler,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) RegisterRoutes() {
	routerGroup := r.router.Group("/example")
	{
		routerGroup.GET("", r.authMiddleware("user"), r.handler.Example)
	}

}
