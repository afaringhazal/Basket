package main

import (
	"log/slog"
	"net/http"
	"os"

	"Basket/handler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	h := handler.Result{
		From:   "Golang",
		Logger: logger.With("handler", "hello"),
	}

	mux := http.NewServeMux()
	print("gggggg")

	mux.HandleFunc("GET /hello", h.Get)
	mux.HandleFunc("POST /hello", h.Post)
	mux.HandleFunc("GET /hello/{username}", h.User)

	if err := http.ListenAndServe("0.0.0.0:1373", mux); err != nil {
		logger.Error("http server failed", "error", err.Error())
	}
}
