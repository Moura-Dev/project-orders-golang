package order_items_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/moura-dev/project-orders-golang/models"
	"github.com/moura-dev/project-orders-golang/repository/order_items_repository"
)

func Create(ctx *gin.Context) {
	var orderItems models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems.Id = int32(itemId)
	orderItems, err = order_items_repository.Create(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func Delete(ctx *gin.Context) {
	orderItemID := ctx.Param("id")
	orderItemIDInt, err := strconv.Atoi(orderItemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID := ctx.Param("productID")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = order_items_repository.Delete(productIDInt, orderItemIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Item deleted successfully")
}

func Get(ctx *gin.Context) {
	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems, err := order_items_repository.Get(orderIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func Update(ctx *gin.Context) {
	var orderItem models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	strOrderId := ctx.Param("id")
	orderId, err := strconv.Atoi(strOrderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItem.OrderId = int32(orderId)
	orderItem, err = order_items_repository.Update(orderItem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItem)
}
