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
			routers.POST("/user", controllers.CreateUser)
			routers.POST("/seller", controllers.CreateSeller)
			routers.GET("/seller", controllers.GetSeller)
			routers.DELETE("/seller", controllers.DeleteSeller)
			routers.GET("/user", controllers.GetUserInfo)
			routers.GET("/company", controllers.GetCompany)
			routers.POST("/company", controllers.CreateCompany)
			routers.DELETE("/company", controllers.DeleteCompany)
			routers.PATCH("/company", controllers.UpdateCompany)
			routers.GET("/product", controllers.GetProduct)
			routers.POST("/product", controllers.CreateProduct)
			routers.PATCH("/product", controllers.UpdateProduct)
			routers.DELETE("/product", controllers.DeleteProduct)
			routers.GET("/address", controllers.GetAddresses)
			routers.POST("/address", controllers.CreateAddress)
			routers.PATCH("/address", controllers.UpdateAddress)
			routers.DELETE("/address", controllers.DeleteAddress)

		}
	}
	return router
}
