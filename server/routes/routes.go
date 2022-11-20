package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/moura-dev/project-orders-golang/controllers/address_controller"
	"github.com/moura-dev/project-orders-golang/controllers/company_controller"
	"github.com/moura-dev/project-orders-golang/controllers/order_items_controller"
	"github.com/moura-dev/project-orders-golang/controllers/orders_controller"
	"github.com/moura-dev/project-orders-golang/controllers/products_controller"
	"github.com/moura-dev/project-orders-golang/controllers/users_controller"
	middlewares "github.com/moura-dev/project-orders-golang/server/middleware"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Authentication"}

	router.Use(cors.New(config))

	main := router.Group("api/")
	main.POST("/login", users_controller.Login)
	main.POST("/user", users_controller.Create)
	{
		routers := main.Group("/", middlewares.AuthJwt())
		{
			// Users
			routers.DELETE("/user/:id", users_controller.Delete)
			routers.GET("/user/:id", users_controller.Get)

			// Companies
			routers.DELETE("/company/:id", company_controller.Delete)
			routers.GET("/company", company_controller.Get)
			//routers.POST("/company", company_controller.Create)
			//routers.PUT("/company/", company_controller.Update)

			// Products
			routers.GET("/product/:id", products_controller.Get)
			routers.DELETE("/product/:id", products_controller.Delete)
			routers.POST("/product", products_controller.Create)
			routers.PUT("/product", products_controller.Update)
			routers.GET("/product", products_controller.GetArray)

			// Address
			routers.DELETE("/address/:id", address_controller.Delete)
			routers.GET("/address", address_controller.Get)
			routers.POST("/address", address_controller.Create)
			routers.PUT("/address", address_controller.Update)

			// Orders
			routers.DELETE("/order/:id", orders_controller.Delete)
			//routers.GET("/order", orders_controller.Get)
			routers.POST("/order", orders_controller.Create)
			routers.PUT("/order", orders_controller.Update)

			// Items
			routers.DELETE("/order/:id/item/:productID", order_items_controller.Delete)
			routers.GET("/order/:id/item", order_items_controller.Get)
			routers.POST("/order/:id/item", order_items_controller.Create)
			routers.PUT("/order/:id/item", order_items_controller.Update)

		}
	}
	return router
}
