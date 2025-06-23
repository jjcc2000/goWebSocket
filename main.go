package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection: %v", err)
		return
	}

	defer conn.Close()

	for {
		// Read message from the client
		messageType, messagge, err := conn.ReadMessage()
		if err != nil {
			fmt.Errorf("Error while reading message: %v", err)
			break
		}
		fmt.Printf("Received: %s\n", messagge)

		// Send the message back to the client
		err = conn.WriteMessage(messageType, messagge)
		if err != nil {
			fmt.Println("Error while writing message :", err)
			break
		}

	}

}

func main() {
	fmt.Println("Websocket is running on..... 8080")
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
