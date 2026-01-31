package getusernotes

import (
	"encoding/json"
	"log/slog"
	resp "mini-blog/internal/models/domain/responce_DTO"
	"mini-blog/pkg/sl"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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

		notes, err := storage.GetUserNotes(userID)
		if err != nil {
			logger.Info("add a description") // TODO
			logger.Error(op, sl.Attr(err))
			return
		}

		marshaled, err := json.Marshal(notes)
		if err != nil {
			logger.Info("internal server error")
			logger.Error(op, sl.Attr(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(marshaled)
		//преобразовать массив структур в JSON, как?

	}
}
