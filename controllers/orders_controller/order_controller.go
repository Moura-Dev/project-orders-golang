package orders_controller

import (
	"base-project-api/models"
	"base-project-api/repository/order_repository"
	"base-project-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	userId, err := services.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order

	if err = ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.UserId = int32(userId)

	order, err = order_repository.Create(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func Delete(ctx *gin.Context) {
	userId, err := services.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID := ctx.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = order_repository.Delete(orderIDInt, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Order deleted successfully")
}

func Update(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]
	userId, err := services.NewJWTService().GetUserIdFromToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.UserId = int32(userIdInt)
	order, err = order_repository.Update(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func Get(ctx *gin.Context) {
	userId, err := services.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := order_repository.Get(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//include total in order
	for i := 0; i < len(orders); i++ {
		total, err := services.SumValues(int(orders[i].Id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orders[i].Total = total
	}

	ctx.JSON(http.StatusOK, orders)
}
