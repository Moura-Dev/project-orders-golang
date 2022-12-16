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
	orderItemID := ctx.Param("id")
	orderItemIDInt, err := strconv.Atoi(orderItemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems.OrderID = int32(orderItemIDInt)
	orderItems, err = repository.InsertItemsOrder(orderItems)
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
	productID := ctx.Param("productID")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.DeleteOrderItems(productIDInt, orderItemIDInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Product deleted successfully")
}

func GetAllItemsInOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems, err := repository.GetAllItemsInOrder(orderIDInt)
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

	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems.OrderID = int32(orderIDInt)

	orderItems, err = repository.UpdateOrderItems(orderItems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItems)
}
