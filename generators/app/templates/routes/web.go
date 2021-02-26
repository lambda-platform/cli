package routes

import (
    "<%= projectName %>/app/controllers"
    "<%= projectName %>/app/middlewares"
    "github.com/labstack/echo/v4"
    "github.com/lambda-platform/agent/agentMW"
)

func Web(e *echo.Echo) {

    //WEB ROUTE
    e.GET("/", controllers.HomeProduction, middlewares.ViewParser())     //production
    e.GET("/dev", controllers.HomeDevelopment, middlewares.ViewParser()) // dev

    //ADMIN ROUTE
    e.GET("/admin", controllers.AdminIndex(true), agentMW.IsLoggedInCookie)
}
