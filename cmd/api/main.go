package main

import (
	"log"
	"net/http"
	"os"
	"suiren/internal/database"
	"suiren/internal/handler"
	"suiren/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Init(os.Getenv("DATABASE_URL"))
	defer db.Close()

	// Rep Handler
	userRepo := repository.NewUserRepo(db)
	userHandler := handler.NewUserHandler(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/users", userHandler.Create)
	r.Get("/users", userHandler.List)
	r.Get("/users/{id}", userHandler.GetByID)

	port := os.Getenv("PORT")
	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
