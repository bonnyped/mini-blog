package getaccesstoken

import (
	"log/slog"
	"mini-blog/pkg/sl"
	"mini-blog/storage/postgres"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Authenticator interface {
	GetMarshaledToken(logger *slog.Logger, userID int) ([]byte, error)
}

func New(logger *slog.Logger, storage postgres.Storage, authenticator Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const (
			op       = "handlers.get_access_token.New"
			paramKey = "id"
		)

		logger = logger.With(slog.String("operation", op),
			slog.String("request id", middleware.GetReqID(r.Context())))

		userID, err := GetIDByKey(paramKey, r)
		if err != nil {
			logger.Info("incorrect value with param, expected integer", "param", paramKey)
			logger.Error(op, sl.Attr(err))
			return
		}

		err = storage.UserExists(userID)
		if err != nil {
			logger.Info("user not exists")
			logger.Error(op, sl.Attr(err))
			http.Error(w, "user not exists", http.StatusNotFound)
		}

		marshaledToken, err := authenticator.GetMarshaledToken(logger, userID)
		if err != nil {
			logger.Info("internal server error") //может нужно более информативно?
			logger.Error("error with marshal token", sl.Attr(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshaledToken)

		logger.Info("access token was issued")
	}
}

func GetIDByKey(key string, r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}
