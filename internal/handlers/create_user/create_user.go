package createuser

import (
	"log/slog"
	"mini-blog/internal/auth"
	req "mini-blog/internal/models/domain/request_DTO"
	"mini-blog/pkg/sl"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

type CreateUser interface {
	CreateUser(username string) (int, error)
}

func New(logger *slog.Logger, userCreator CreateUser, jwtManager auth.JWTManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create_user.New"

		logger = logger.With(slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var user req.User

		user.Username = r.FormValue("username")
		if user.Username == "" {
			http.Error(w, "username is required", http.StatusBadRequest)
			return
		}

		id, err := userCreator.CreateUser(user.Username)
		if err != nil {
			logger.Error(op, sl.Attr(err))
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		marshaledToken, err := jwtManager.GetMarshaledToken(logger, id)
		if err != nil {
			//TODO delete created user
			logger.Error("failed to get access token", sl.Attr(err))
			http.Error(w, "failed to create user, internal error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshaledToken)

		logger.Info("User created", "user name", user.Username)
	}
}
