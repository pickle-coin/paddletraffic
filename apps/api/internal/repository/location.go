package repository

import (
	"paddletraffic/internal/model"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) Create(location *model.Location) error {
	return r.db.Create(location).Error
}

func (r *LocationRepository) GetByID(id uint) (*model.Location, error) {
	var location model.Location
	err := r.db.First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepository) GetAll() ([]model.Location, error) {
	var locations []model.Location
	err := r.db.Find(&locations).Error
	return locations, err
}

func (r *LocationRepository) Update(location *model.Location) error {
	return r.db.Save(location).Error
}

func (r *LocationRepository) Delete(id uint) error {
	return r.db.Delete(&model.Location{}, id).Error
}
