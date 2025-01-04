package repositories

import (
	"context"

	"github.com/anderson2093/apiGolang/domain/models"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByID(ctx context.Context, id string) (*models.Category, error)
	Delete(ctx context.Context, id string) error
}
