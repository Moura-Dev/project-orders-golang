package routes

import (
	"base-project-api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		routers := main.Group("/")
		{
			routers.GET("/", controllers.HellowControllers)
		}
	}
	return router
}
