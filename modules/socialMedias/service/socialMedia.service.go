package service

import (
	"errors"
	repoPhoto "final_project_hacktiv8/modules/photos/repository"
	"final_project_hacktiv8/modules/socialMedias/dto"
	"final_project_hacktiv8/modules/socialMedias/model"
	repoSocialmedia "final_project_hacktiv8/modules/socialMedias/repository"
	"final_project_hacktiv8/modules/socialMedias/validation"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ServiceSocialMedia interface {
	Create(data dto.Request) (dto.Response, error)
	GetList() (dto.ResponseListWrapper, error)
	UpdateByID(data dto.Request) (dto.Response, error)
	DeleteByID(id uint) error
}

type service struct {
	repository repoSocialmedia.RepositorySocialMedia
	repoPhoto  repoPhoto.RepositoryPhoto
}

func New(repository repoSocialmedia.RepositorySocialMedia, repoPhoto repoPhoto.RepositoryPhoto) ServiceSocialMedia {
	return &service{repository: repository, repoPhoto: repoPhoto}
}

func (srv *service) Create(data dto.Request) (dto.Response, error) {
	err := validation.ValidateSocialMediaCreate(data)
	if err != nil {
		return dto.Response{}, err
	}

	socialMedia := new(model.SocialMedia)
	copier.Copy(&socialMedia, &data)

	createdSocialMedia, err := srv.repository.Create(*socialMedia)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{}
	copier.Copy(&response, &createdSocialMedia)

	return response, nil
}

func (srv *service) GetList() (dto.ResponseListWrapper, error) {
	listSocialMedia, err := srv.repository.GetList()
	if err != nil {
		return dto.ResponseListWrapper{}, err
	}

	responseList := []dto.ResponseList{}

	for _, socialMedia := range listSocialMedia {
		//get photo by user id
		photo, err := srv.repoPhoto.GetPhotoByUserID(socialMedia.UserID)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ResponseListWrapper{}, err
		}

		response := new(dto.ResponseList)
		copier.Copy(&response, &socialMedia)
		if photo.PhotoURL != "" {
			response.User.ProfileImageUrl = photo.PhotoURL
		}

		responseList = append(responseList, *response)
	}

	return dto.ResponseListWrapper{SocialMedias: responseList}, nil
}

func (srv *service) UpdateByID(data dto.Request) (dto.Response, error) {
	err := validation.ValidateSocialMediaCreate(data)
	if err != nil {
		return dto.Response{}, err
	}

	socialMedia := new(model.SocialMedia)
	copier.Copy(&socialMedia, &data)

	updatedSocialMedia, err := srv.repository.UpdateByID(*socialMedia)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{}
	copier.Copy(&response, &updatedSocialMedia)

	return response, nil
}

func (srv *service) DeleteByID(id uint) error {
	err := srv.repository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
