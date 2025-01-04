package routes

import (
	"github.com/anderson2093/apiGolang/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, categoryController *controllers.CategoryController) {
	categoryGroup := router.Group("/categories")
	{
		// @Summary Crear una nueva categoría
		// @Description Crea una nueva categoría en la tienda virtual
		// @Tags Categorías
		// @Accept json
		// @Produce json
		// @Param category body models.Category true "Detalles de la categoría"
		// @Success 201 {object} models.Category
		// @Failure 400 {object} models.ErrorResponse
		// @Failure 500 {object} models.ErrorResponse
		// @Router /categories [post]
		categoryGroup.POST("/", categoryController.CreateCategory)

		// @Summary Obtener todas las categorías
		// @Description Obtener todas las categorías disponibles en la tienda
		// @Tags Categorías
		// @Accept json
		// @Produce json
		// @Success 200 {array} models.Category
		// @Failure 400 {object} models.ErrorResponse
		// @Router /categories [get]
		categoryGroup.GET("/", categoryController.GetCategories)
		// Similar para las demás rutas
	}
}
