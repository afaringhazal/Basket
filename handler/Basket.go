package handler

import (
	"Basket/database"
	"Basket/model"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type ConsId struct {
	Val int
}

type Result_3 struct {
	From   string
	Logger *slog.Logger
}

func (h Result_3) AddBasket(w http.ResponseWriter, r *http.Request) {

	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bas model.Basket

	if err := json.NewDecoder(r.Body).Decode(&bas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	print(bas.Username)
	print(bas.Data)
	print(bas.State)

	if !validationUser(bas.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id_ := GenerateId()

	originalBas := model.Basket{
		Id:         id_(),
		Username:   bas.Username,
		Data:       bas.Data,
		State:      bas.State,
		Created_at: time.Now(),
		Update_at:  time.Now(),
	}

	result := database.DB.Create(originalBas) // pass pointer of data to Create
	print(result)
	w.WriteHeader(http.StatusNoContent)
	return
}

func validationUser(username string) bool {
	var user model.User
	database.DB.Where("Username = ?", username).First(&user)
	if user.Username == "" {
		return false
	}
	return true
}

func GenerateId() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
func (h Result_3) GetAllBaskets(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("get all basket!")
	username := r.PathValue("username")
	var baskets []model.Basket
	result := database.DB.Where("Username = ?", username).Find(&baskets)
	if result.Error != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusOK)
	json_baskets, _ := json.Marshal(baskets)

	w.Write(json_baskets)
	return
}

func (h Result_3) UpdateBasket(w http.ResponseWriter, r *http.Request) {

	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bas model.Basket

	if err := json.NewDecoder(r.Body).Decode(&bas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	print(bas.Username)
	print(bas.Data)
	print(bas.State)

	username := bas.Username

	basketID := bas.Id

	var existingBasket model.Basket
	if err := database.DB.Where(" Username = ? AND id = ?", username, basketID).First(&existingBasket).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if existingBasket.State == true { // that is mean complete
		w.WriteHeader(http.StatusForbidden)
	}

	existingBasket.Data = bas.Data
	existingBasket.State = bas.State

	database.DB.Save(&existingBasket)
	w.WriteHeader(http.StatusOK)
	return
}

func (h Result_3) Delete(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("delete basket!")
	username := r.PathValue("username")
	basketID := r.PathValue("id")
	var basket model.Basket
	if err := database.DB.Where("Username = ? AND id = ?", username, basketID).Delete(&basket).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	return
}
