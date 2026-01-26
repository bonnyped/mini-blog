package main

import (
	"log/slog"
	"mini-blog/internal/config"
	createnote "mini-blog/internal/handlers/create_note"
	createuser "mini-blog/internal/handlers/create_user"
	"mini-blog/storage/postgres"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting server...")

	config := config.MustLoad(logger)

	r := chi.NewRouter()

	logger.Info("Setting up middleware...")
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)

	storage, err := postgres.New(logger, config.DbServer)
	if err != nil {
		logger.Error("Failed to initialize storage", "error", err)
		return
	}

	r.Post("/users", createuser.New(logger, storage, config.JWTSecret))
	r.Post("/users/{id}/notes", createnote.New(logger, storage))

	logger.Info("Listening on 127.0.0.1:8082")

	err = http.ListenAndServe("127.0.0.1:8082", r)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
	}
	//TODO: config implementation

	//TODO: storage implementation

	//TODO: logger implementation

}
