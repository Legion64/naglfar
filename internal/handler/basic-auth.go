package handler

import (
	"auth-app/internal/crypto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Credentials struct {
		Username string `json:"username,omitempty" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	RegisterCredentials struct {
		Credentials
		Email   string `json:"email" validate:"required,email"`
		Name    string `json:"name" validate:"required,min=2,max=32"`
		Surname string `json:"surname" validate:"required,min=2,max=32"`
	}
)

func LoginHandler(c echo.Context) error {
	var credentials Credentials

	if err := c.Bind(&credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, credentials)
}

func RegisterHandler(c echo.Context) error {
	var rc RegisterCredentials

	if err := c.Bind(&rc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&rc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := crypto.HashPassword(rc.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User Registered",
	})
}
