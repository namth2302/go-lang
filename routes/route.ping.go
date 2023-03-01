package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namth/go-examples/handlers/ping"
)

func InitPingRoutes(route *gin.Engine) {
	pingHandler := handlers.NewHandlerPing()
	groupRoute := route.Group("/ping")

	groupRoute.GET("", pingHandler.PingHandler)
}
