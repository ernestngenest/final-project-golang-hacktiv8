package repository

import (
	"final_project_hacktiv8/modules/socialMedias/model"

	"gorm.io/gorm"
)

type RepositorySocialMedia interface {
	Create(data model.SocialMedia) (model.SocialMedia, error)
	GetList() ([]model.SocialMedia, error)
	UpdateByID(data model.SocialMedia) (model.SocialMedia, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositorySocialMedia {
	return &repository{db: db}
}

func (r *repository) Create(data model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return data, nil
}

func (r *repository) GetList() ([]model.SocialMedia, error) {
	var socialMedia []model.SocialMedia
	err := r.db.Preload("User").Find(&socialMedia).Error
	if err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (r *repository) UpdateByID(data model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Model(&data).Updates(&data).First(&data).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return data, nil
}

func (r *repository) DeleteByID(id uint) error {
	socialMedia := new(model.SocialMedia)
	socialMedia.ID = id
	return r.db.First(&socialMedia).Where("id = ?", socialMedia.ID).Delete(&socialMedia).Error
}
