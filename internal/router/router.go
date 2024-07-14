package router

import (
	"auth-app/internal/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "All systems are operational!")
	})

	v1 := e.Group("/v1/auth")
	{
		v1.POST("", handler.LoginHandler)
		v1.POST("/register", handler.RegisterHandler)
	}
}
