package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(any interface{}) error {
	err := validate.Struct(any)

	return err
}
