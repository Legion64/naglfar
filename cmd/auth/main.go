package main

import (
	"auth-app/internal/db"
	"auth-app/internal/router"
	"auth-app/internal/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.RegisterRouter(e)
	validator.RegisterValidator(e)

	if dbConn, err := db.New(); err == nil {
		dbConn.Migrate()
	}

	e.Logger.Fatal(e.Start(":8080"))
}
