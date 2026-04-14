package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"golang-crud/src/config"
	"golang-crud/src/models"
	"golang-crud/src/utils"
	"golang-crud/src/validation"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

	if err := validation.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)


	config.DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	config.DB.Where("email = ?", input.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		fmt.Println("Error generating token", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
