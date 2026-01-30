package createnote

import (
	"log/slog"
	domain "mini-blog/internal/models/domain/request_DTO"
	"mini-blog/pkg/sl"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type CreateNote interface {
	CreateNote(userId uint64, note domain.Note) error
}

func New(logger *slog.Logger, noteCreator CreateNote) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create_note.New"

		logger = logger.With(slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var note domain.Note

		userId, err := strconv.Atoi((chi.URLParam(r, "id")))
		if err != nil {
			logger.Error(op, "error", "incorrect user_id")
			return
		}

		err = render.DecodeJSON(r.Body, &note)
		if err != nil {
			logger.Error(op, sl.Attr(err))
			return
		}

		err = noteCreator.CreateNote(uint64(userId), note)
		if err != nil {
			logger.Error("failed to create note", sl.Attr(err))
			return
		}
	}
}
