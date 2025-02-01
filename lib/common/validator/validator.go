package validator

import (
	"common/validation"
	"sync"

	"github.com/go-playground/validator/v10"
)

// Глобальный валидатор
var (
	vinstance *validator.Validate
	once      sync.Once
)

func GetValidatorInstance() *validator.Validate {
	once.Do(func() {
		vinstance = validator.New()
		vinstance.RegisterValidation("cdnURL", validation.CdnURLValidation)
		vinstance.RegisterValidation("dateonly", validation.DateOnlyValidation)
		vinstance.RegisterValidation("minage", validation.MinAgeValidation)
	})
	return vinstance
}
