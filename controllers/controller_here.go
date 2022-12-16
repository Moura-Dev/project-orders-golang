package controllers

import "github.com/gin-gonic/gin"

func (c *Controller) HellowControllers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hello World!",
	})
}
