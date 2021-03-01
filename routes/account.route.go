package routes

import (
	"github.com/TelegramServer/controllers"
	"github.com/TelegramServer/middlewares"
)

// InitAccountRouter inits account router
func InitAccountRouter() {
	middlewares.AccountRouter.HandleFunc("/login", controllers.HandleLogin).Methods("GET")
}
