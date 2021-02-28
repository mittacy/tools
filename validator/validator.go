package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidatorStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// 传入的值不是结构体等无效值时
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("you may have to use the ValidatorVar()")
		}
		msg := ""
		for _ , err := range err.(validator.ValidationErrors) {
			msg += err.Namespace() + " rules: " + err.Tag() + " " + err.Param() + "\n"
		}
		return errors.New(msg)
	}
	return nil
}

func ValidatorVar(s interface{}, tagName, tagRule string) error {
	err := validate.Var(s, tagRule)
	if err != nil {
		msg := tagName
		for _ , err := range err.(validator.ValidationErrors) {
			msg += " rules: " + err.Tag() + " " + err.Param()
		}
		return errors.New(msg)
	}
	return nil
}