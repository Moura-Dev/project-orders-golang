package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := repository.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func UpdateProduct(ctx *gin.Context) {
	var product models.Product
	//productID := ctx.Param("id")
	//productIdInt, err := strconv.Atoi(productID)

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := repository.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productID)

	err = repository.DeleteProduct(productIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Product deleted successfully")
}

func GetProductByID(ctx *gin.Context) {
	var product models.Product
	productID := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productID)

	product, err = repository.GetProductByID(productIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func GetAllProducts(ctx *gin.Context) {
	products, err := repository.GetAllProduts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
