package controllers

import (
	"net/http"

	"github.com/TelegramServer/middlewares"
	"github.com/TelegramServer/models"
	"github.com/TelegramServer/utils"
)

// HandleLogin use to handle logic request from client
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body
	phone := r.Form.Get("phone")
	password := r.Form.Get("password")
	user := models.GetUserModel().GetUserByPhone(phone)

	if user == nil {
		utils.SendJSON(w, false)
	}

	tokenString, err := utils.GenerateTokenString(user.Phone)
	if err != nil {
		middlewares.ErrorHandler(err, w, r, http.StatusNotFound)
		return
	}

	if user.Password == password {
		utils.SendJSON(w, tokenString)
		return
	}

}
