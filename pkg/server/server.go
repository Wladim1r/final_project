package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func Run(r *chi.Mux) error {
	webDir := "./web"

	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	port := "7540"
	if p := os.Getenv("TODO_PORT"); p != "" {
		port = p
	}

	log.Printf("Приложение запущено на порту: %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}

	return nil
}
