package api

import (
	"github.com/TampelliniOtavio/backend-web-development-course/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSUpportedCurrency(currency)
	}

	return false
}
