package controllers

import (
	"context"
	"net/http"

	"github.com/anderson2093/apiGolang/domain/models"
	"github.com/anderson2093/apiGolang/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	usecase *usecase.CategoryUsecase
}

func NewCategoryController(u *usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{usecase: u}
}

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
func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.usecase.CreateCategory(context.Background(), &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// @Summary Obtener todas las categorías
// @Description Obtener todas las categorías disponibles en la tienda
// @Tags Categorías
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Router /categories [get]
func (ctrl *CategoryController) GetCategories(c *gin.Context) {
	categories, err := ctrl.usecase.GetCategories(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// Similar para GetCategoryByID y DeleteCategory...
