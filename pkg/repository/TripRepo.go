package repository

import (
	"OrderDelayServing/pkg/model"
	"context"
	"gorm.io/gorm"
)

type TripRepo interface {
	Create(ctx context.Context, trip model.Trip) error
	Get(ctx context.Context, tripID uint) (model.Trip, error)
	Update(ctx context.Context, trip model.Trip) error
	Delete(ctx context.Context, tripID uint) error
}

type TripRepoImpl struct {
	db *gorm.DB
}

func NewTripRepo(db *gorm.DB) *TripRepoImpl {
	return &TripRepoImpl{
		db: db,
	}
}

func (r *TripRepoImpl) Create(ctx context.Context, trip model.Trip) error {
	result := r.db.WithContext(ctx).Create(trip)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TripRepoImpl) Get(ctx context.Context, tripID uint) (model.Trip, error) {
	var trip model.Trip
	result := r.db.WithContext(ctx).Where(&model.Trip{ID: tripID}).First(&trip)
	if result.Error != nil {
		return model.Trip{}, result.Error
	}

	return trip, nil
}

func (r *TripRepoImpl) Update(ctx context.Context, trip model.Trip) error {
	result := r.db.WithContext(ctx).Where(&model.Trip{ID: trip.ID}).Updates(&trip)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TripRepoImpl) Delete(ctx context.Context, tripID uint) error {
	result := r.db.WithContext(ctx).Where(&model.Trip{ID: tripID}).Delete(&model.Trip{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
