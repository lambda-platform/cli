package routes

import (
	"github.com/labstack/echo/v4"
	"<%= projectName %>/<%= serviceName %>/handlers"
)

func Api(e *echo.Echo) {
	a := e.Group("/api")

	a.GET("/users", handlers.Users)

}

