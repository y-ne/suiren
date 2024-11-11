package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/rs/zerolog/log"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) HelloWorld(c echo.Context) error {
    log.Info().Msg("Received request to root endpoint")
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Hello, World!",
    })
}