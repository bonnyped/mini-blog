package getusernotes

import (
	"log/slog"
	resp "mini-blog/internal/models/domain/responce_DTO"
	"mini-blog/pkg/sl"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var notes []resp.UserNote

type AllUserNotesHandler interface {
	GetUserNotes(userID int) ([]resp.UserNote, error)
}

func New(logger *slog.Logger, storage AllUserNotesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.get_user_notes.New"

		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logger.Info("incorrect format of UserID")
			logger.Error(op, sl.Attr(err))
			return
		}

		notes, err = storage.GetUserNotes(userID)
		if err != nil {
			logger.Info("add a description") // TODO
			logger.Error(op, sl.Attr(err))
			return
		}
		_ = notes

		w.Header().Set("status", "200 - OK")
		//преобразовать массив структур в JSON, как?

	}
}
