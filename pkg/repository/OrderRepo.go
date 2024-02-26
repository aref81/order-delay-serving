package repository

import (
	"OrderDelayServing/pkg/model"
	"context"
	"gorm.io/gorm"
)

type OrderRepo interface {
	Create(ctx context.Context, order model.Order) (model.Order, error)
	Get(ctx context.Context, orderID uint) (model.Order, error)
	GetOrdersByVendorID(ctx context.Context, vendorID uint) ([]model.Order, error)
	Update(ctx context.Context, order model.Order) error
	Delete(ctx context.Context, orderID uint) error
}

type OrderRepoImpl struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepoImpl {
	return &OrderRepoImpl{
		db: db,
	}
}

func (r *OrderRepoImpl) Create(ctx context.Context, order model.Order) (model.Order, error) {
	result := r.db.WithContext(ctx).Create(&order)
	if result.Error != nil {
		return model.Order{}, result.Error
	}

	return order, nil
}

func (r *OrderRepoImpl) Get(ctx context.Context, orderID uint) (model.Order, error) {
	var order model.Order
	result := r.db.WithContext(ctx).Where(&model.Order{ID: orderID}).First(&order)
	if result.Error != nil {
		return model.Order{}, result.Error
	}

	return order, nil
}

func (r *OrderRepoImpl) GetOrdersByVendorID(ctx context.Context, vendorID uint) ([]model.Order, error) {
	var orders []model.Order
	result := r.db.WithContext(ctx).Where(&model.Order{VendorID: vendorID}).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (r *OrderRepoImpl) Update(ctx context.Context, order model.Order) error {
	result := r.db.WithContext(ctx).Where(&model.Order{ID: order.ID}).Updates(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepoImpl) Delete(ctx context.Context, orderID uint) error {
	result := r.db.WithContext(ctx).Where(&model.Order{ID: orderID}).Delete(&model.Order{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
