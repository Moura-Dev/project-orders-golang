package controllers

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	CreateSeller(ctx *gin.Context)
	HellowControllers(ctx *gin.Context)
	GetSellerById(ctx *gin.Context)
	GetAllSellers(ctx *gin.Context)
	//UpdateSeller(ctx *gin.Context)
	//DeleteSeller(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	GetAllCompany(ctx *gin.Context)
	//CreateCompany(ctx *gin.Context)
	//DeleteCompany(ctx *gin.Context)
	//UpdateCompany(ctx *gin.Context)
	//GetAllProducts(ctx *gin.Context)
	//GetProductByID(ctx *gin.Context)
	//CreateProduct(ctx *gin.Context)
	//UpdateProduct(ctx *gin.Context)
	//DeleteProduct(ctx *gin.Context)
	//GetAllAddress(ctx *gin.Context)
	//CreateAddress(ctx *gin.Context)
	//UpdateAddress(ctx *gin.Context)
	//DeleteAddress(ctx *gin.Context)
	//GetAllOrdersByID(ctx *gin.Context)
	//CreateOrder(ctx *gin.Context)
	//UpdateOrder(ctx *gin.Context)
	//DeleteOrder(ctx *gin.Context)
	//InsertItemsOrder(ctx *gin.Context)
	//UpdateOrderItems(ctx *gin.Context)
	//DeleteOrderItems(ctx *gin.Context)
	//GetAllItemsInOrder(ctx *gin.Context)
}

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

type MockController struct {
}

func NewMockController() MockController {
	return MockController{}
}
