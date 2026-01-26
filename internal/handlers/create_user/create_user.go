package createuser

import (
	"log/slog"
	"mini-blog/internal/models/domain"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

type CreateUser interface {
	CreateUser(logger slog.Logger, username string, secret string) error
}

func New(logger *slog.Logger, userCreator CreateUser, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create_user.New"

		logger = logger.With(slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var user domain.User

		user.Username = r.FormValue("name")
		if user.Username == "" {
			http.Error(w, "username is required", http.StatusBadRequest)
			return
		}

		err := userCreator.CreateUser(*logger, user.Username, secret)
		if err != nil {
			logger.Error("failed to create user", slog.String("error", err.Error()))
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
	}
}
