package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *restErrors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return restErrors.NewBadRequestError("Invalid Field Type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorCauses := []restErrors.Cause{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := restErrors.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return restErrors.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return restErrors.NewBadRequestError("Error trying to convert fields")
	}
}
