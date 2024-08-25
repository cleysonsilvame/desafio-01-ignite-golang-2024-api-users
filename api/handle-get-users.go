package api

import (
	"net/http"
)

func handleGetUsers(app application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := app.findAll()

		sendJSON(
			w,
			Response{Data: users},
			http.StatusOK,
		)
	}
}
