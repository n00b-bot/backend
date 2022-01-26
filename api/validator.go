package api

import (
	"github.com/go-playground/validator/v10"
)

const (
	USD = "USD"
	VND = "VND"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, VND:
		return true
	default:
		return false
	}
}

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return IsSupportedCurrency(currency)
	}
	return false
}
