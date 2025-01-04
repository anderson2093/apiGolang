package usecase

import (
	"context"

	"github.com/anderson2093/apiGolang/domain/models"
	"github.com/anderson2093/apiGolang/domain/repositories"
)

type CategoryUsecase struct {
	repo repositories.CategoryRepository
}

func NewCategoryUsecase(repo repositories.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) CreateCategory(ctx context.Context, category *models.Category) error {
	return u.repo.Create(ctx, category)
}

func (u *CategoryUsecase) GetCategories(ctx context.Context) ([]models.Category, error) {
	return u.repo.GetAll(ctx)
}

func (u *CategoryUsecase) GetCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *CategoryUsecase) DeleteCategory(ctx context.Context, id string) error {
	return u.DeleteCategory(ctx, id)
}
