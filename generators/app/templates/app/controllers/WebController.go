package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "github.com/foolin/goview/supports/echoview-v4"
)

//home page
func HomeProduction(c echo.Context) error {
    return echoview.Render(c, http.StatusOK, "home_production", map[string]interface{}{})
}

func HomeDevelopment(c echo.Context) error {
    return echoview.Render(c, http.StatusOK, "home_development", map[string]interface{}{})
}

func HomeData() map[string]interface{} {

    return map[string]interface{}{
        "title": "test",
    }
}
