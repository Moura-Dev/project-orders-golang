package controllers

//
//import (
//	"base-project-api/models"
//	"base-project-api/repository"
//	"base-project-api/services"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//"base-project-api/models"
//	"base-project-api/repository"
//	"base-project-api/services"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//func CreateOrder(ctx *gin.Context) {
//	var order models.Order
//
//	if err := ctx.ShouldBindJSON(&order); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	order.UserID = int32(userIdInt)
//
//	order, err = repository.CreateOrder(order)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, order)
//}
//
//func DeleteOrder(ctx *gin.Context) {
//
//	orderID := ctx.Param("id")
//	orderIDInt, err := strconv.Atoi(orderID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	err = repository.DeleteOrder(orderIDInt, userIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, "Order deleted successfully")
//}
//
//func UpdateOrder(ctx *gin.Context) {
//	var order models.Order
//
//	if err := ctx.ShouldBindJSON(&order); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	order.UserID = int32(userIdInt)
//	order, err = repository.UpdateOrder(order)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, order)
//}
//
//func GetAllOrdersByID(ctx *gin.Context) {
//	token := ctx.Request.Header.Get("authorization")
//	token = token[7:]
//	userId, err := services.NewJWTService().GetUserIdFromToken(token)
//	if err != nil {
//		ctx.JSON(401, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//	userIdInt, err := strconv.Atoi(userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	orders, err := repository.GetAllOrders(userIdInt)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	//include total in order
//	for i := 0; i < len(orders); i++ {
//		total, err := services.SumValues(int(orders[i].ID))
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		orders[i].Total = total
//	}
//
//	ctx.JSON(http.StatusOK, orders)
//}
