package api

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"api-users/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func handleUpdateUserByID(app application) http.HandlerFunc {
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

		r.Body = http.MaxBytesReader(w, r.Body, 1024*1)

		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxErr *http.MaxBytesError

			if errors.As(err, &maxErr) {
				slog.Error("request body to large", "error", err)
				sendJSON(
					w,
					Response{Error: "request body to large"},
					http.StatusBadRequest,
				)
				return
			}

			slog.Error("error reading request body", "error", err)
			sendJSON(
				w,
				Response{Error: "something went wrong"},
				http.StatusInternalServerError,
			)
			return
		}

		var payload struct {
			Name string `json:"name"`
			Bio  string `json:"bio"`
		}

		if err := json.Unmarshal(data, &payload); err != nil {
			slog.Error("error on unmarshal user", "error", err)
			sendJSON(w, Response{Error: "invalid json"}, http.StatusBadRequest)
			return
		}

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

		user.
			app.updateByID(user)

		if valid, _ := user.ValidateUser(); !valid {
			sendJSON(
				w,
				Response{Error: "Please provided first_name, last_name and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		id := app.insert(user)

		sendJSON(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}
