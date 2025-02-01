package validation

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

func MinAgeValidation(fl validator.FieldLevel) bool {
	birthDateStr := fl.Field().String()
	minAge := fl.Param()

	requiredAge, err := strconv.Atoi(minAge)
	if err != nil {
		return false
	}

	birthDate, err := time.Parse(time.DateOnly, birthDateStr)
	if err != nil {
		return false
	}

	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age >= requiredAge
}
