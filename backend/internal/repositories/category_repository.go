package repositories

import (
	"context"

	"github.com/PrinceNarteh/pos/internal/models"
	"gorm.io/gorm"
)

var _ CategoryRepository = (*categoryRepository)(nil)

type CategoryRepository interface {
	FindAll(context.Context) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func (r *categoryRepository) categoryTbl() gorm.Interface[models.Category] {
	return gorm.G[models.Category](r.db)
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	return r.categoryTbl().Find(ctx)
}

// func (r *categoryRepository) Create(ctx context.Context, createCategoryDTO dto.CreateCategoryDTO) {
// 	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
// 	defer cancel()
// 	err := r.categoryTbl().Create(ctx,, crecreateCategoryDTO)
// }
