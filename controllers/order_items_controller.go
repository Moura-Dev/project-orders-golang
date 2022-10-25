package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InsertItemsOrder(ctx *gin.Context) {
	var orderItems models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems, err := repository.InsertItemsOrder(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func DeleteOrderItems(ctx *gin.Context) {
	orderItemID := ctx.Param("id")
	orderItemIDInt, err := strconv.Atoi(orderItemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order_id := 3

	err = repository.DeleteOrderItems(orderItemIDInt, order_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Item deleted successfully")
}

func GetAllItemsInOrder(ctx *gin.Context) {
	order_id := 4
	orderItems, err := repository.GetAllItemsInOrder(order_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func UpdateOrderItems(ctx *gin.Context) {
	var orderItems models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems, err := repository.UpdateOrderItems(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}
