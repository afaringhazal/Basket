package main

import (
	"Basket/database"
	"log/slog"
	"net/http"
	"os"

	"Basket/handler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	U := handler.Result_2{
		From:   "User",
		Logger: logger.With("handler", "User"),
	}

	B := handler.Result_3{
		From:   "Basket",
		Logger: logger.With("handler", "Basket"),
	}

	database.Connect()

	mux := http.NewServeMux()
	// user
	mux.HandleFunc("POST /user/{username}/{password}", U.AddUser) // (returns a list of baskets)
	mux.HandleFunc("GET  /user/login/{username}/{password}", U.Login)
	mux.HandleFunc("GET /user/{username}", U.GetUser) // (returns a list of baskets)

	//basket
	mux.HandleFunc("POST /basket", B.AddBasket)
	mux.HandleFunc("GET /basket/{username}", B.GetAllBaskets)
	mux.HandleFunc("PATCH /basket/{id}", B.UpdateBasket) //updates the given basket)
	mux.HandleFunc("DELETE /basket/{id}", B.Delete)      //(deletes the given backset)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		logger.Error("http server failed", "error", err.Error())
	}

}
