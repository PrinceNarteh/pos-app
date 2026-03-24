package repositories

import (
	"context"

	"github.com/PrinceNarteh/pos/internal/dto"
	"github.com/PrinceNarteh/pos/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ CartRepository = (*cartRepository)(nil)

type CartRepository interface {
	FindAll(ctx context.Context) ([]models.Cart, error)
	FindByID(ctx context.Context, cartID string) (*models.Cart, error)
	Create(ctx context.Context, data *models.Cart) error
	Update(ctx context.Context, cartID string, updates *dto.UpdateCartDTO) (*models.Cart, error)
	Delete(ctx context.Context, cardID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func (r *cartRepository) cartTbl() gorm.Interface[models.Cart] {
	return gorm.G[models.Cart](r.db)
}

func (r *cartRepository) FindAll(ctx context.Context) ([]models.Cart, error) {
	return r.cartTbl().Find(ctx)
}

func (r *cartRepository) FindByID(ctx context.Context, cartID string) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	cart, err := r.cartTbl().Where("id = ?", cartID).First(ctx)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, gorm.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &cart, nil
}

func (r *cartRepository) Create(ctx context.Context, data *models.Cart) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	return r.cartTbl().Create(ctx, data)
}

func (r *cartRepository) Update(ctx context.Context, cartID string, updates *dto.UpdateCartDTO) (*models.Cart, error) {
	cart := new(models.Cart)
	res := r.db.Model(cart).
		Clauses(clause.Returning{}).
		Where("id = ?", cartID).
		Updates(updates)

	if res.Error != nil {
		return nil, res.Error
	}

	return cart, nil
}

func (r *cartRepository) Delete(ctx context.Context, cartID string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rowsAffected, err := r.cartTbl().Where("id = ?", cartID).Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
