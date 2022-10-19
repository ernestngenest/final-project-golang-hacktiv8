package service

import (
	"final_project_hacktiv8/modules/comments/dto"
	"final_project_hacktiv8/modules/comments/model"
	"final_project_hacktiv8/modules/comments/repository"
	"final_project_hacktiv8/modules/comments/validation"
	"log"

	"github.com/jinzhu/copier"
)

type ServiceComment interface {
	Create(request dto.Request) (dto.ResponseInsert, error)
	Get() ([]dto.Response, error)
	Update(request dto.RequestUpdate, commentID uint) (dto.ResponseUpdate, error)
	Delete(commentID uint) error
}

type service struct {
	repo repository.RepositoryComment
}

func New(repo repository.RepositoryComment) ServiceComment {
	return &service{repo: repo}
}

func (s *service) Create(request dto.Request) (dto.ResponseInsert, error) {
	err := validation.ValidateComment(request)
	if err != nil {
		return dto.ResponseInsert{}, err
	}

	var comment model.Comment
	copier.Copy(&comment, &request)
	create, err := s.repo.Create(comment)
	if err != nil {
		return dto.ResponseInsert{}, err
	}
	var response dto.ResponseInsert
	copier.Copy(&response, &create)
	return response, nil
}

func (s *service) Get() ([]dto.Response, error) {
	comments, err := s.repo.Get()
	if err != nil {
		return []dto.Response{}, err
	}

	var response []dto.Response

	for _, comment := range comments {
		var singleResponse dto.Response
		copier.Copy(&singleResponse, &comment)
		response = append(response, singleResponse)
	}

	return response, nil
}

func (s *service) Update(request dto.RequestUpdate, commentID uint) (dto.ResponseUpdate, error) {
	// validate request
	err := validation.ValidateCommentUpdate(request)
	if err != nil {
		return dto.ResponseUpdate{}, err
	}
	// update db with repo
	var comment model.Comment
	copier.Copy(&comment, request)
	comment.ID = commentID
	update, err := s.repo.Update(comment)
	if err != nil {
		return dto.ResponseUpdate{}, err
	}
	log.Println(update)
	var responseComment dto.ResponseUpdate
	copier.Copy(&responseComment, update)
	return responseComment, nil
}

func (s *service) Delete(commentID uint) error {
	err := s.repo.Delete(commentID)
	if err != nil {
		return err
	}
	return nil
}
