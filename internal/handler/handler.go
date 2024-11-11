// internal/handler/handler.go
package handler

import (
	"database/sql"
	"net/http"
	"suiren/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(dbConn *sql.DB) *Handler {
	return &Handler{
		queries: db.New(dbConn),
	}
}

func (h *Handler) HelloWorld(c echo.Context) error {
	log.Info().Msg("Received request to root endpoint")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}

func (h *Handler) ListUsers(c echo.Context) error {
	log.Info().Msg("Listing all users")

	users, err := h.queries.ListUsers(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("Failed to list users")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get users")
	}

	return c.JSON(http.StatusOK, users)
}
