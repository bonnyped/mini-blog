package createuser

import (
	"log/slog"
	"mini-blog/internal/models/domain"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

type CreateUser interface {
	CreateUser(username string, creationTime time.Time) error
}

func New(log *slog.Logger, userCreator CreateUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create_user.New"

		log = log.With(slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var user domain.User

		user.Username = r.FormValue("name")
		if user.Username == "" {
			http.Error(w, "username is required", http.StatusBadRequest)
			return
		}

		err := userCreator.CreateUser(user.Username, time.Now())
		if err != nil {
			log.Error("failed to create user", slog.String("error", err.Error()))
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
	}
}
