package connections

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Connection *websocket.Conn

func Subscribe(conn *websocket.Conn) {
	Connection = conn
}

func SendMessage(message string) {
	if Connection != nil {
		fmt.Println("Pushing to web socket")
		err := Connection.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error during message writing:", err)
		}
	}
}
