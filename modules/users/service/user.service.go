package service

// import user.dto.go

import (
	"final_project_hacktiv8/global"
	"final_project_hacktiv8/helpers"
	"final_project_hacktiv8/modules/users/dto"
	"final_project_hacktiv8/modules/users/model"
	"final_project_hacktiv8/modules/users/repository"
	"final_project_hacktiv8/modules/users/validation"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	Create(data dto.Request) (dto.Response, error)
	Login(data dto.RequestLogin) (dto.ResponseLogin, error)
	Update(data dto.Request) (dto.Response, error)
	DeleteByID(id uint) error
}

type service struct {
	repo repository.RepositoryUser
}

func New(repo repository.RepositoryUser) ServiceUser {
	return &service{repo: repo}
}

func (s *service) Create(data dto.Request) (dto.Response, error) {
	err := validation.ValidateUserCreate(data, s.repo)
	if err != nil {
		return dto.Response{}, err
	}

	User := new(model.User)

	copier.Copy(&User, &data)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.Response{}, err
	}
	User.Password = string(hashedPassword)

	createdUser, err := s.repo.Create(*User)
	if err != nil {
		return dto.Response{}, err
	}

	resp := dto.Response{}

	copier.Copy(&resp, &createdUser)
	resp.UpdatedAt = nil

	return resp, nil
}

func (s *service) Login(data dto.RequestLogin) (dto.ResponseLogin, error) {
	err := validation.ValidateUserLogin(data)
	if err != nil {
		return dto.ResponseLogin{}, err
	}

	dataUser, err := s.repo.Login(data.Email)
	if err != nil {
		return dto.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password))
	if err != nil {
		return dto.ResponseLogin{}, global.ErrorInvalidLogin
	}

	token, err := helpers.NewJwt(dataUser.ID)
	if err != nil {
		return dto.ResponseLogin{}, err
	}

	resp := dto.ResponseLogin{}
	resp.Token = token

	return resp, nil
}

func (s *service) Update(data dto.Request) (dto.Response, error) {
	err := validation.ValidateUserUpdate(data)
	if err != nil {
		return dto.Response{}, err
	}

	user := model.User{}
	copier.Copy(&user, &data)

	updatedUser, err := s.repo.Update(user)
	if err != nil {
		return dto.Response{}, err
	}

	resp := dto.Response{}

	copier.Copy(&resp, &updatedUser)

	return resp, nil
}

func (s *service) DeleteByID(id uint) error {
	return s.repo.DeleteByID(id)
}
