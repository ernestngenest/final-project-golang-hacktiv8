package validation

import (
	"final_project_hacktiv8/modules/photos/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidatePhotoCreate(data dto.Request) error {
	return validation.Errors{
		"title":     validation.Validate(data.Title, validation.Required),
		"caption":   validation.Validate(data.Caption),
		"photo_url": validation.Validate(data.PhotoURL, validation.Required),
		"userId":    validation.Validate(data.UserID),
	}.Filter()
}
