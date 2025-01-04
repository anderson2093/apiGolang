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

func (ctrl *CategoryController) GetCategories(c *gin.Context) {
	categories, err := ctrl.usecase.GetCategories(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// Similar para GetCategoryByID y DeleteCategory...
