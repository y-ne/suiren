package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Home(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}
}
