package service

import (
	"final_project_hacktiv8/modules/photos/dto"
	"final_project_hacktiv8/modules/photos/model"
	"final_project_hacktiv8/modules/photos/repository"
	"final_project_hacktiv8/modules/photos/validation"

	"github.com/jinzhu/copier"
)

type ServicePhoto interface {
	Create(data dto.Request) (dto.Response, error)
	GetPhotos() ([]dto.ResponseGet, error)
	Update(data dto.Request, photoID int) (dto.ResponseUpdate, error)
	Delete(photoID int) error
}

func New(photoRepo repository.RepositoryPhoto) ServicePhoto {
	return &service{RepositoryPhoto: photoRepo}
}

type service struct {
	RepositoryPhoto repository.RepositoryPhoto
}

// Create photo to DB
func (service *service) Create(data dto.Request) (dto.Response, error) {
	// validation input
	err := validation.ValidatePhotoCreate(data)
	if err != nil {
		return dto.Response{}, err
	}

	photo := new(model.Photo)

	copier.Copy(&photo, &data)

	create, err := service.RepositoryPhoto.Create(*photo)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{}
	copier.Copy(&response, &create)

	return response, nil
}

// Update photo from DB
func (service *service) Update(data dto.Request, photoID int) (dto.ResponseUpdate, error) {
	// validate update request
	err := validation.ValidatePhotoCreate(data)
	if err != nil {
		return dto.ResponseUpdate{}, err
	}

	photo := model.Photo{}
	copier.Copy(&photo, &data)
	photo.ID = uint(photoID)

	// call repository method to update Photo
	update, err := service.RepositoryPhoto.Update(photo)
	if err != nil {
		return dto.ResponseUpdate{}, err
	}
	resp := dto.ResponseUpdate{}
	copier.Copy(&resp, &update)
	return resp, nil
}

// Get all photos from DB
func (service *service) GetPhotos() ([]dto.ResponseGet, error) {
	resPhotos, err := service.RepositoryPhoto.GetPhotos()

	if err != nil {
		return []dto.ResponseGet{}, nil
	}

	var response []dto.ResponseGet
	for _, photo := range resPhotos {
		tempResp := dto.ResponseGet{}
		tempResp.ID = photo.ID
		tempResp.Title = photo.Title
		tempResp.Caption = photo.Caption
		tempResp.PhotoURL = photo.PhotoURL
		tempResp.CreatedAt = photo.CreatedAt
		tempResp.UpdatedAt = photo.UpdatedAt
		tempResp.User.Username = photo.User.Username
		tempResp.User.Email = photo.User.Email
		response = append(response, tempResp)
	}

	return response, nil
}

// Delete photo from DB by photo ID
func (service *service) Delete(photoID int) error {
	err := service.RepositoryPhoto.Delete(photoID)
	if err != nil {
		return err
	}
	return nil
}
