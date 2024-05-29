package handler

import (
	"Basket/model"
	"encoding/json"
	"log/slog"
	"net/http"
)
import "Basket/database"

type Result_2 struct {
	From   string
	Logger *slog.Logger
}

func (h Result_2) AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	password := r.PathValue("password")

	h.Logger.Info("read username from path parameter", "username", username, " password:", password)

	user := model.User{username, password}
	user.HashPassword(user.Password)

	result := database.DB.Create(user) // pass pointer of data to Create
	print(result)
	w.WriteHeader(http.StatusNoContent)
	return
}

func (h Result_2) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	h.Logger.Info("read username from path parameter", "username", username)

	var user model.User
	database.DB.Where("username = ?", username).First(&user)

	print("****\n")

	//result, _ := database.DB.Get(username) // pass pointer of data to Create
	print(user.Username)
	print(user.Password)
	w.WriteHeader(http.StatusOK)
	json_user, _ := json.Marshal(user)

	w.Write(json_user)
	return
}
