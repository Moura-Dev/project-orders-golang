package routes

import (
	"base-project-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Authentication"}

	router.Use(cors.New(config))

	mockController := controllers.NewMockController()

	controller := controllers.NewController()

	main := router.Group("api/")
	main.POST("/login", controllers.Login)
	main.POST("/user", controllers.CreateUser)
	{
		routers := main.Group("/")
		routers.Use(cors.New(config))
		{
			routers.GET("/", controller.HellowControllers)
			routers.POST("/seller", mockController.CreateSeller)
			routers.GET("/seller/:id", mockController.GetSellerById)
			routers.GET("/seller", mockController.GetAllSellers)
			routers.DELETE("/seller/:id", controllers.DeleteSeller)
			routers.GET("/user", controllers.GetUserInfo)
			routers.GET("/company", mockController.GetAllCompany)
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
