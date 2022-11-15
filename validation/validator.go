package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateStruct(data any) []*ErrorResponse {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Message = msgForTag(err)
			errors = append(errors, &element)
		}
	}
	return errors
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error()
}
