package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgradation:", err)
		return
	}
	fmt.Println("Socket connected")

	// err = conn.WriteMessage(websocket.TextMessage, []byte("Let's start to talk something."))
	// if err != nil {
	// 	fmt.Println("Error during message writing:", err)
	// }

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error during message reading:", err)
			break
		}
		fmt.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, []byte("Let's start to talk something."))
		if err != nil {
			fmt.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func main() {
	fmt.Println("Setting up server in 8081")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/hey", home)
	http.ListenAndServe(":8081", nil)
}
