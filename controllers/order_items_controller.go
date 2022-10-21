package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderItems(ctx *gin.Context) {
	var orderItems models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems, err := repository.CreateOrderItems(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func DeleteOrderItems(ctx *gin.Context) {
	var orderItems models.OrderItem

	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems, err := repository.DeleteOrderItems(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}

func GetOrderItems(ctx *gin.Context) {
	var items models.OrderItem
	orderItems, err := repository.GetOrderItems(items)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}
