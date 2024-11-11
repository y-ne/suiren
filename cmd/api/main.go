package main

import (
    "os"
    "time"
    "suiren/internal/handler"

    "github.com/labstack/echo/v4"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    // Setup zerolog
    output := zerolog.ConsoleWriter{
        Out:        os.Stdout,
        TimeFormat: time.RFC3339,
    }
    log.Logger = zerolog.New(output).With().Timestamp().Logger()

    // Create Echo instance
    e := echo.New()
    e.HideBanner = true

    // Initialize handler
    h := handler.NewHandler()

    // Routes
    e.GET("/", h.HelloWorld)

    // Start server
    port := ":8080"
    log.Info().Msgf("Server starting on http://localhost%s", port)
    if err := e.Start(port); err != nil {
        log.Fatal().Err(err).Msg("Failed to start server")
    }
}