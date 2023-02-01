package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppStatus struct {
	Status string `json:"status" example:"UP"`
}

func Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK, AppStatus{Status: "UP"})
}
