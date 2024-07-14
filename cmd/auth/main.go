package main

import (
	"github.com/Legion64/naglfar/internal/db"
	"github.com/Legion64/naglfar/internal/router"
	"github.com/Legion64/naglfar/internal/validator"
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
