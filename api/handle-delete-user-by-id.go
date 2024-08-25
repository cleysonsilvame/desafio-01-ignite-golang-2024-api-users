package api

import (
	"log/slog"
	"net/http"

	"api-users/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func handleDeleteUserByID(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		parsedID, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("error on parsing uuid from path params", "error", err)
			sendJSON(
				w,
				Response{
					Error: "error on parsing id",
				},
				http.StatusBadRequest,
			)
			return
		}

		id := domain.ID(parsedID)

		ok := app.deleteByID(id)

		if !ok {
			sendJSON(
				w,
				Response{
					Error: "user not found",
				},
				http.StatusNotFound,
			)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
