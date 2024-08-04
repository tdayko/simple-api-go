package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create opening"})
}

func ShowOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Show opening"})
}

func DeleteOpeningByIdOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete opening"})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update opening"})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List opening"})
}