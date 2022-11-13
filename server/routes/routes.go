package routes

import (
	"base-project-api/controllers"
	middlewares "base-project-api/server/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/", middlewares.CORSMiddleware())
	main.POST("/login", controllers.Login)
	main.POST("/user", controllers.CreateUser)
	{
		routers := main.Group("/", middlewares.AuthJwt(), middlewares.CORSMiddleware())
		{
			routers.GET("/", controllers.HellowControllers)
			routers.POST("/seller", controllers.CreateSeller)
			routers.GET("/seller/:id", controllers.GetSellerById)
			routers.GET("/seller", controllers.GetAllSellers)
			routers.DELETE("/seller/:id", controllers.DeleteSeller)
			routers.GET("/user", controllers.GetUserInfo)
			routers.GET("/company", controllers.GetAllCompany)
			routers.POST("/company", controllers.CreateCompany)
			routers.DELETE("/company/:id", controllers.DeleteCompany)
			routers.PUT("/company/", controllers.UpdateCompany)
			routers.GET("/product", controllers.GetAllProducts)
			routers.GET("/product/:id", controllers.GetProductByID)
			routers.POST("/product", controllers.CreateProduct)
			routers.PUT("/product", controllers.UpdateProduct)
			routers.DELETE("/product/:id", controllers.DeleteProduct)
			routers.GET("/address", controllers.GetAllAddress)
			routers.POST("/address", controllers.CreateAddress)
			routers.PUT("/address", controllers.UpdateAddress)
			routers.DELETE("/address/:id", controllers.DeleteAddress)
			routers.GET("/order", controllers.GetAllOrdersByID)
			routers.POST("/order", controllers.CreateOrder)
			routers.PUT("/order", controllers.UpdateOrder)
			routers.DELETE("/order/:id", controllers.DeleteOrder)
			routers.POST("/order/:id/item", controllers.InsertItemsOrder)
			routers.PUT("/order/:id/item", controllers.UpdateOrderItems)
			routers.DELETE("/order/:id/item/:productID", controllers.DeleteOrderItems)
			routers.GET("/order/:id/item", controllers.GetAllItemsInOrder)

		}
	}
	return router
}
