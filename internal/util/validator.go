package util

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
)

var validate = validator.New()

func BindAndValidate(c echo.Context, req any) error {
	// Bind request
	if err := c.Bind(req); err != nil {
		return errz.NewValidation("bad request body")
	}

	// Validate request
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := validate.Struct(v.Index(i).Interface()); err != nil {
				return errz.NewValidation(err.Error())
			}
		}
	default:
		if err := validate.Struct(req); err != nil {
			return errz.NewValidation(err.Error())
		}
	}

	return nil
}
