package repository

import (
	"OrderDelayServing/pkg/model"
	"context"
	"gorm.io/gorm"
)

type DelayReportRepo interface {
	Create(ctx context.Context, report model.DelayReport) (model.DelayReport, error)
	Get(ctx context.Context, reportID uint) (model.DelayReport, error)
	Update(ctx context.Context, report model.DelayReport) error
	Delete(ctx context.Context, reportID uint) error
}

type DelayReportRepoImpl struct {
	db *gorm.DB
}

func NewDelayReportRepo(db *gorm.DB) *DelayReportRepoImpl {
	return &DelayReportRepoImpl{
		db: db,
	}
}

func (r *DelayReportRepoImpl) Create(ctx context.Context, report model.DelayReport) (model.DelayReport, error) {
	result := r.db.WithContext(ctx).Create(&report)
	if result.Error != nil {
		return model.DelayReport{}, result.Error
	}

	return report, nil
}

func (r *DelayReportRepoImpl) Get(ctx context.Context, reportID uint) (model.DelayReport, error) {
	var report model.DelayReport
	result := r.db.WithContext(ctx).Where(&model.DelayReport{ID: reportID}).First(&report)
	if result.Error != nil {
		return model.DelayReport{}, result.Error
	}

	return report, nil
}

func (r *DelayReportRepoImpl) Update(ctx context.Context, report model.DelayReport) error {
	result := r.db.WithContext(ctx).Where(&model.DelayReport{ID: report.ID}).Updates(&report)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *DelayReportRepoImpl) Delete(ctx context.Context, reportID uint) error {
	result := r.db.WithContext(ctx).Where(&model.DelayReport{ID: reportID}).Delete(&model.DelayReport{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
