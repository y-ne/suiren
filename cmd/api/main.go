// cmd/api/main.go
package main

import (
	"database/sql"
	"os"
	"suiren/internal/handler"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load()

	// Setup zerolog
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	// Connect to database using database/sql
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	// Configure the connection pool
	db.SetMaxOpenConns(25)                 // Maximum number of open connections
	db.SetMaxIdleConns(10)                 // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Unable to ping database")
	}

	// Create Echo instance
	e := echo.New()
	e.HideBanner = true

	// Initialize handler
	h := handler.NewHandler(db)

	// Routes
	e.GET("/", h.HelloWorld)
	e.GET("/users", h.ListUsers)

	// Start server
	port := ":8080"
	log.Info().Msgf("Server starting on http://localhost%s", port)
	if err := e.Start(port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
