package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MainValidator struct {
	validator *validator.Validate
}

func RegisterValidator(e *echo.Echo) {
	e.Validator = &MainValidator{
		validator: validator.New(),
	}
}

func (m MainValidator) Validate(i interface{}) error {
	if err := m.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
