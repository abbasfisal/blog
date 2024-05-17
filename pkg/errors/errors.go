package errors

import (
	"errors"
	"github.com/abbasfisal/blog/internal/providers/validation"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errorList = make(map[string]string)

func Init() {
	errorList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			Add(fieldErr.Field(), GetErrorMessage(fieldErr.Tag()))
		}
	}
}
func Add(key string, value string) {
	errorList[strings.ToLower(key)] = value
}

func GetErrorMessage(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorList
}
