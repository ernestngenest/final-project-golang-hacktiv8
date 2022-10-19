package repository

import (
	"final_project_hacktiv8/modules/photos/model"

	"gorm.io/gorm"
)

type RepositoryPhoto interface {
	Create(data model.Photo) (model.Photo, error)
	GetPhotos() ([]model.Photo, error)
	Update(data model.Photo) (model.Photo, error)
	Delete(id int) error
	GetPhotoByUserID(id uint) (model.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryPhoto {
	return &repository{db: db}
}

// Create data Photo
func (r *repository) Create(data model.Photo) (model.Photo, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.Photo{}, err
	}
	return data, nil
}

// GetPhotos  return slice photo
func (r *repository) GetPhotos() ([]model.Photo, error) {
	var photo []model.Photo
	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return []model.Photo{}, err
	}
	return photo, nil
}

// Update photo from DB
func (r *repository) Update(data model.Photo) (model.Photo, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.Photo{}, err
	}
	return data, nil
}

// Delete photo from DB by photo ID
func (r *repository) Delete(id int) error {
	photo := model.Photo{}
	photo.ID = uint(id)
	err := r.db.First(&photo).Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

// GetPhoto By User ID
func (r *repository) GetPhotoByUserID(id uint) (model.Photo, error) {
	var photo model.Photo
	err := r.db.Preload("User").Where("user_id = ?", id).First(&photo).Error
	if err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}
