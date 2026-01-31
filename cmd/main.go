package main

import (
	"log/slog"
	"mini-blog/internal/config"
	newNote "mini-blog/internal/handlers/create_note"
	newUser "mini-blog/internal/handlers/create_user"
	getaccesstoken "mini-blog/internal/handlers/get_access_token"
	userNotes "mini-blog/internal/handlers/get_user_notes"
	mymiddleware "mini-blog/internal/my_middleware"
	"mini-blog/pkg/sl"
	"mini-blog/storage/postgres"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting server...")

	config := config.MustLoad(logger)

	router := chi.NewRouter()

	logger.Info("Setting up middleware...")

	storage, err := postgres.New(config.DbServer)
	if err != nil {
		logger.Error("Failed to initialize storage", sl.Attr(err))
		return
	}
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.JWTManager.JWTAuth))
		r.Use(jwtauth.Authenticator(config.JWTManager.JWTAuth))
		r.Use(mymiddleware.Authorize)
		r.Post("/users/{id}/notes", newNote.New(logger, storage))
		r.Get("/users/{id}/notes", userNotes.New(logger, storage))
	})

	router.Post("/users", newUser.New(logger, storage, config.JWTManager))
	router.Get("/users/{id}", getaccesstoken.New(logger, *storage, config.JWTManager))

	logger.Info("Listening on 127.0.0.1:8082")

	err = http.ListenAndServe("127.0.0.1:8082", router)
	if err != nil {
		logger.Error("Server failed to start", sl.Attr(err))
	}
}
