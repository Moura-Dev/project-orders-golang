package products_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/moura-dev/project-orders-golang/models"
	"github.com/moura-dev/project-orders-golang/repository/product_repository"
	"github.com/moura-dev/project-orders-golang/services"
)

func Get(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	var product models.Product

	productID := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productID)

	product, err = product_repository.Get(userId, productIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func GetArray(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	products, err := product_repository.GetArray(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func Create(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.UserId = int32(userId)

	product, err := product_repository.Create(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func Update(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.UserId = int32(userId)
	product, err := product_repository.Update(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func Delete(ctx *gin.Context) {
	userId := services.GetUserIdFromContext(ctx)

	strProductId := ctx.Param("id")
	productId, err := strconv.Atoi(strProductId)

	err = product_repository.Delete(userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Product deleted successfully")
}
