package main

import (
	"log/slog"
	"mini-blog/internal/config"
	createuser "mini-blog/internal/handlers/create_user"
	"mini-blog/storage/postgres"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("Starting server...")
	config := config.MustLoad()
	log.Info("Config loaded successfully...")
	r := chi.NewRouter()
	log.Info("Starting chi router...")
	log.Info("Setting up middleware...")
	r.Use(middleware.RequestID)
	//TODO: настроить конфиг для всех переменных

	storage, err := postgres.New(config.DBCOnnectionString)
	if err != nil {
		log.Error("Failed to initialize storage", "error", err)
		return
	}

	r.Post("/users", createuser.New(log, storage))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request", "method", r.Method, "path", r.URL.Path)
		w.Write([]byte("Hello, gopher!"))
	})

	log.Info("Listening on 127.0.0.1:8082")

	err = http.ListenAndServe("127.0.0.1:8082", r)
	if err != nil {
		log.Error("Server failed to start", "error", err)
	}
	//TODO: config implementation

	//TODO: storage implementation

	//TODO: logger implementation

}
