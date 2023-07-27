package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techshool/simplebank/util"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		//check currency if supported
		return util.IsSupportCurrency(currency)
	}
	return false
}
