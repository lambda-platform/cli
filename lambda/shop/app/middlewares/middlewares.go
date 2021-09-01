package middlewares

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func Set(e *echo.Echo) {
    /*
    	|----------------------------------------------
    	| Echo useful MIDDLEWARES
    	|----------------------------------------------
    */
    e.Use(middleware.Secure())
    e.Use(middleware.Recover())
    //App.Echo.Use(middleware.Logger())e
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins:     []string{"*", "http://localhost:*", "http://127.0.0.1:*"},
        AllowCredentials: true,
        AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "x-requested-with"},
        AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
    }))
}
