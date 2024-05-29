package handler

import (
	"Basket/request"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Result struct {
	From   string
	Logger *slog.Logger
}

func (h Result) User(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("username")

	h.Logger.Info("read username from path parameter", "username", value)

	w.WriteHeader(http.StatusNoContent)
}

func (h Result) Get(w http.ResponseWriter, r *http.Request) {
	if value := r.FormValue("hello"); value != "" {
		h.Logger.Info("read hello from query parameter", "hello", value)
	}

	enc, err := json.Marshal(fmt.Sprintf("Result World from %s", h.From))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}

func (h Result) Post(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req request.Name

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if req.Count == nil {
		h.Logger.Info("There is no count")
	} else {
		h.Logger.Info("There is a count", "count", *req.Count)
	}

	enc, err := json.Marshal(fmt.Sprintf("Result to %s from %s", req.Name, h.From))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}
