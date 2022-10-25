package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := repository.CreateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	userId := 1
	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repository.DeleteOrder(orderIDInt, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Order deleted successfully")
}

func UpdateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := repository.UpdateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func GetOrders(ctx *gin.Context) {
	userID := 1
	orders, err := repository.GetAllOrders(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
