package handlers

import (
	"fmt"
	"net/http"

	"consumer/app/connections"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	connections.Subscribe(conn)
	if err != nil {
		fmt.Println("Error during connection upgradation:", err)
		return
	}
	fmt.Println("Web Socket connected")

	defer conn.Close()

	for {
	}
}
