package controllers

import (
    "<%= projectName %>/app/models"
    "github.com/labstack/echo/v4"
    "github.com/lambda-platform/lambda/DB"
    "net/http"
)

func Users(c echo.Context) error {

    users := []models.Users{}
    DB.DB.Find(&users)

    return c.JSON(http.StatusOK, users)
}
