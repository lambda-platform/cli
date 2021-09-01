package routes

import (
    "lambda/app/controllers"
    "github.com/labstack/echo/v4"
)

func Api(e *echo.Echo) {
    a := e.Group("/api")
    a.GET("/users", controllers.Users)
}
