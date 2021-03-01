package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TelegramServer/middlewares"
	"github.com/TelegramServer/models"
	"github.com/TelegramServer/routes"
	"github.com/TelegramServer/utils"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Message is a struct
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		middlewares.ErrorHandler(err, w, r, http.StatusForbidden)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func initRouters() {
	middlewares.Router.HandleFunc("/ws", handleConnections)
	routes.InitAccountRouter()
}

func main() {
	disconnect := utils.ConnectDatabase("telegram")
	defer disconnect()
	user := models.GetUserModel().GetUserByPhone("0963612265")
	fmt.Println(user.Name)

	initRouters()
	http.Handle("/", middlewares.Router)
	go handleMessages()

	log.Println("start server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
