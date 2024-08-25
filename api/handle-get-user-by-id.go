package api

import (
	"log/slog"
	"net/http"

	"api-users/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func handleGetUserByID(app application) http.HandlerFunc {
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

		user := app.findByID(id)

		if user == nil {
			sendJSON(
				w,
				Response{
					Error: "user not found",
				},
				http.StatusNotFound,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}
