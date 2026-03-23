package repositories

import (
	"context"
	"errors"

	"github.com/PrinceNarteh/pos/internal/dto"
	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ SupplierRepository = (*supplierRepository)(nil)

type SupplierRepository interface {
	FindAll(ctx context.Context) ([]models.Supplier, error)
	FindByID(ctx context.Context, supplierID string) (*models.Supplier, error)
	Create(ctx context.Context, data *models.Supplier) error
	Update(ctx context.Context, supplierID string, update *dto.UpdateSupplierDTO) (*models.Supplier, error)
	Delete(ctx context.Context, supplierID string) error
}

type supplierRepository struct {
	db *gorm.DB
}

func (r *supplierRepository) supplierTbl() gorm.Interface[models.Supplier] {
	return gorm.G[models.Supplier](r.db)
}

func (r *supplierRepository) FindAll(ctx context.Context) ([]models.Supplier, error) {
	return nil, nil
}

func (r *supplierRepository) FindByID(ctx context.Context, supplierID string) (*models.Supplier, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	supplier, err := r.supplierTbl().Where("id = ?", supplierID).First(ctx)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, gorm.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &supplier, nil
}

func (r *supplierRepository) Create(ctx context.Context, data *models.Supplier) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := r.supplierTbl().Create(ctx, data)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "suppliers_email_key" {
				return ErrDuplicateEmail
			}
		}
	}

	return nil
}

func (r *supplierRepository) Update(ctx context.Context, supplierID string, updates *dto.UpdateSupplierDTO) (*models.Supplier, error) {
	supplier := new(models.Supplier)

	res := r.db.Model(supplier).
		Clauses(clause.Returning{}).
		Where("id = ?", supplierID).
		Updates(updates)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return supplier, nil
}

func (r *supplierRepository) Delete(ctx context.Context, supplierID string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rowsAffected, err := r.supplierTbl().Where("id = ?", supplierID).Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
