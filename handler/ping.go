package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Ping(c echo.Context) error {
	err := h.service.Ping()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "pong")
}
