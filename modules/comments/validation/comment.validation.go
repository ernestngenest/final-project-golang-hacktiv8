package validation

import (
	"final_project_hacktiv8/modules/comments/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateComment(data dto.Request) error {
	return validation.Errors{
		"message":  validation.Validate(data.Message, validation.Required),
		"photo_id": validation.Validate(data.PhotoID, validation.Required),
	}.Filter()
}

func ValidateCommentUpdate(data dto.RequestUpdate) error {
	return validation.Errors{
		"message": validation.Validate(data.Message, validation.Required),
	}.Filter()
}
