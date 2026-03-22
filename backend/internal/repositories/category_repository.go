package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_                CategoryRepository = (*categoryRepository)(nil)
	ErrDuplicateName                    = errors.New("category name already exists")
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, categoryID string) (*models.Category, error)
	FindByName(ctx context.Context, name string) (*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, categoryID string, name string) (*models.Category, error)
	Delete(ctx context.Context, categoryID string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func (r *categoryRepository) categoryTbl() gorm.Interface[models.Category] {
	return gorm.G[models.Category](r.db)
}

func (r *categoryRepository) findBy(ctx context.Context, query, value string) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	category, err := r.categoryTbl().Where(query, value).First(ctx)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, gorm.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &category, err
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	return r.categoryTbl().Find(ctx)
}

func (r *categoryRepository) FindByID(ctx context.Context, categoryID string) (*models.Category, error) {
	return r.findBy(ctx, "id = ?", categoryID)
}

func (r *categoryRepository) FindByName(ctx context.Context, name string) (*models.Category, error) {
	return r.findBy(ctx, "name = ?", name)
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if err := r.categoryTbl().Create(ctx, category); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "categories_name_key" {
				return ErrDuplicateName
			}
		}
		return fmt.Errorf("error creating category: %w", err)
	}

	return nil
}

func (r *categoryRepository) Update(ctx context.Context, categoryID string, name string) (*models.Category, error) {
	category := new(models.Category)
	res := r.db.Model(category).
		Clauses(clause.Returning{}).
		Where("id = ?", categoryID).
		Update("name", name)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return category, nil
}

func (r *categoryRepository) Delete(ctx context.Context, categoryID string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rowsAffected, err := r.categoryTbl().Where("id = ?", categoryID).Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
