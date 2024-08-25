package api

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"api-users/domain"
)

func handleCreateUser(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		var user domain.User

		if err := json.Unmarshal(data, &user); err != nil {
			slog.Error("error on unmarshal user", "error", err)
			sendJSON(w, Response{Error: "invalid json"}, http.StatusBadRequest)
			return
		}

		if valid, _ := user.ValidateUser(); !valid {
			sendJSON(
				w,
				Response{Error: "Please provided first_name, last_name and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		id := app.insert(user)

		sendJSON(w, Response{Data: id.String()}, http.StatusOK)
	}
}
