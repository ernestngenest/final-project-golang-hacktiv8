package repository

import (
	"errors"

	"final_project_hacktiv8/global"
	"final_project_hacktiv8/modules/users/model"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(data model.User) (model.User, error)
	IsEmailExist(email string) error
	Login(email string) (model.User, error)
	Update(data model.User) (model.User, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryUser {
	return &repository{db: db}
}

func (r *repository) Create(data model.User) (model.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}

func (r *repository) IsEmailExist(email string) error {
	user := new(model.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	return global.ErrorEmailAlreadyExists
}

func (r *repository) Login(email string) (model.User, error) {
	user := new(model.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func (r *repository) Update(data model.User) (model.User, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}

func (r *repository) DeleteByID(id uint) error {
	user := new(model.User)
	user.ID = id
	return r.db.First(&user).Where("id = ?", user.ID).Delete(&user).Error
}
