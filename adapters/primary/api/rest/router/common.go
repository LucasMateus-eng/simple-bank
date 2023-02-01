package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppStatus struct {
	Status string `json:"status" example:"UP"`
}

// Health handler to check if API is available
//
// @Summary Get API availability
// @Description Get API availability - if it's running
// @Tags health
// @Produce json
// @Success 200 {object} AppStatus "API is available."
// @Router /health [get]
func Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK, AppStatus{Status: "UP"})
}
