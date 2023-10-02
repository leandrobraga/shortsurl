package data

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Url struct {
	ID        string    `json:"id"`
	Url       string    `validate:"required,url" json:"url"`
	Code      string    `validate:"required,len=6" json:"code"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *Url) IsValid() (bool, []ErrorResponse) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	var errors []ErrorResponse

	errs := validate.Struct(u)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			})

		}
		return false, errors
	}
	return true, errors
}
