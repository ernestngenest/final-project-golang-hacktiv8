package validation

import (
	"errors"

	"final_project_hacktiv8/modules/users/dto"
	"final_project_hacktiv8/modules/users/repository"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func isEmailExist(repo repository.RepositoryUser) validation.RuleFunc {
	return func(value interface{}) error {
		email, ok := value.(string)
		if !ok {
			return errors.New("invalid email address")
		}

		return repo.IsEmailExist(email)
	}
}

func ValidateUserCreate(data dto.Request, repo repository.RepositoryUser) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email, validation.By(isEmailExist(repo))),
		"username": validation.Validate(data.Username, validation.Required),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20)),
		"age":      validation.Validate(data.Age, validation.Required),
	}.Filter()
}

func ValidateUserLogin(data dto.RequestLogin) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20).Error("invalid email or password")),
	}.Filter()
}

func ValidateUserUpdate(data dto.Request) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"username": validation.Validate(data.Username, validation.Required),
	}.Filter()
}
