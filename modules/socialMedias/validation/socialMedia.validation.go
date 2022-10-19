package validation

import (
	"final_project_hacktiv8/modules/socialMedias/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateSocialMediaCreate(data dto.Request) error {
	return validation.Errors{
		"name":             validation.Validate(data.Name, validation.Required),
		"social_media_url": validation.Validate(data.SocialMediaUrl, validation.Required),
	}.Filter()
}
