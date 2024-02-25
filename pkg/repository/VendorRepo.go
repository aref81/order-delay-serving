package repository

import (
	"OrderDelayServing/pkg/model"
	"context"
	"gorm.io/gorm"
)

type VendorRepo interface {
	Create(ctx context.Context, vendor model.Vendor) error
	Get(ctx context.Context, vendorID uint) (model.Vendor, error)
	Update(ctx context.Context, vendor model.Vendor) error
	Delete(ctx context.Context, vendorID uint) error
}

type VendorRepoImpl struct {
	db *gorm.DB
}

func NewVendorRepo(db *gorm.DB) *VendorRepoImpl {
	return &VendorRepoImpl{
		db: db,
	}
}

func (r *VendorRepoImpl) Create(ctx context.Context, vendor model.Vendor) error {
	result := r.db.WithContext(ctx).Create(vendor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *VendorRepoImpl) Get(ctx context.Context, vendorID uint) (model.Vendor, error) {
	var vendor model.Vendor
	result := r.db.WithContext(ctx).Where(&model.Vendor{ID: vendorID}).First(&vendor)
	if result.Error != nil {
		return model.Vendor{}, result.Error
	}

	return vendor, nil
}

func (r *VendorRepoImpl) Update(ctx context.Context, vendor model.Vendor) error {
	result := r.db.WithContext(ctx).Where(&model.Vendor{ID: vendor.ID}).Updates(&vendor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *VendorRepoImpl) Delete(ctx context.Context, vendorID uint) error {
	result := r.db.WithContext(ctx).Where(&model.Vendor{ID: vendorID}).Delete(&model.Vendor{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
