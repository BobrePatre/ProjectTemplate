package example

import (
	"github.com/BobrePatre/ProjectTemplate/internal/delivery/http/handlers/example"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router         *gin.RouterGroup
	handler        *example.Handler
	authMiddleware webAuthProvider.AuthHttpMiddleware
}

func NewRouter(router *gin.RouterGroup, authMiddleware webAuthProvider.AuthHttpMiddleware, handler *example.Handler) *Router {

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
