package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, r Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(r)
	if err != nil {
		slog.Error("error marshalling response", "error", err)
		sendJSON(
			w,
			Response{
				Error: "something went wrong",
			},
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error writing response", "error", err)
		return
	}
}
