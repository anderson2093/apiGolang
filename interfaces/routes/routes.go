package routes

import (
	"github.com/anderson2093/apiGolang/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, categoryController *controllers.CategoryController) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.POST("/", categoryController.CreateCategory)
		categoryGroup.GET("/", categoryController.GetCategories)
		// Similar para las demás rutas
	}
}
