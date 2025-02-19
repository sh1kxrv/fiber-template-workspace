package validation

import "github.com/go-playground/validator/v10"

const cdnUrl = "<CDN_URL>"

func CdnURLValidation(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	return len(url) >= 22 && url[:22] == cdnUrl
}
