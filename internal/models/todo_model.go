package models

import (
	"github.com/go-playground/validator/v10"
)

type TodoInput struct {
	Description string `json:"description" validate:"required,min=3,max=20"`
}
type TodoUpdate struct {
	Description string `json:"description" validate:"required,min=3,max=20"`
	Completed   string `json:"completed" validate:"required,boolean"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
