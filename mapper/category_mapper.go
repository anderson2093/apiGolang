package mapper

import (
	"github.com/anderson2093/apiGolang/domain/models"
	"github.com/anderson2093/apiGolang/dto/request"
	"github.com/anderson2093/apiGolang/dto/response"
)

func ToCategoryResponse(category models.Category) response.CategoryResponseDto {
	return response.CategoryResponseDto{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoryListResponse(categories []models.Category) []response.CategoryResponseDto {
	var categoryListResponse []response.CategoryResponseDto
	for _, category := range categories {
		categoryListResponse = append(categoryListResponse, ToCategoryResponse(category))
	}
	return categoryListResponse
}

func ToCategory(categoryDto request.CategoryRequestDto) models.Category {
	return models.Category{
		Name:        categoryDto.Name,
		Description: categoryDto.Description,
	}
}
