package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func DateOnlyValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.DateOnly, fl.Field().String())
	return err == nil
}
