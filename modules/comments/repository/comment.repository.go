package repository

import (
	"final_project_hacktiv8/modules/comments/model"

	"gorm.io/gorm"
)

type RepositoryComment interface {
	Create(data model.Comment) (model.Comment, error)
	Get() ([]model.Comment, error)
	Update(data model.Comment) (model.Comment, error)
	Delete(commentID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(data model.Comment) (model.Comment, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.Comment{}, err
	}
	return data, nil
}

func (r repository) Get() ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}

func (r *repository) Update(data model.Comment) (model.Comment, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.Comment{}, err
	}
	return data, nil
}

// Delete comment by comment id return a error or nil
func (r repository) Delete(commentID uint) error {
	var comment model.Comment
	comment.ID = commentID
	err := r.db.First(&comment).Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func New(db *gorm.DB) RepositoryComment {
	return &repository{db: db}
}
