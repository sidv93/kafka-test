package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	kafka "github.com/segmentio/kafka-go"
)

var upgrader = websocket.Upgrader{}

const (
	topic         = "systemstats"
	brokerAddress = "broker:29092"
)

var sConnection *websocket.Conn

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	sConnection = conn
	if err != nil {
		fmt.Println("Error during connection upgradation:", err)
		return
	}
	fmt.Println("Socket connected")

	defer conn.Close()

	for {
		// err = conn.WriteMessage(websocket.TextMessage, []byte("Let's start to talk something."))
		// if err != nil {
		// 	fmt.Println("Error during message writing:", err)
		// 	break
		// }
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func setupKafka(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		if sConnection != nil {
			fmt.Println("sending to socket")
			err = sConnection.WriteMessage(websocket.TextMessage, []byte(string(msg.Value)))
			if err != nil {
				fmt.Println("Error during message writing:", err)
				break
			}
		}
	}
}

func main() {
	fmt.Println("Setting up server in 8081")
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/hey", home)
	ctx := context.Background()
	go setupKafka(ctx)
	http.ListenAndServe(":8081", nil)
}
