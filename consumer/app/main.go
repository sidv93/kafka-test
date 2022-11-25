package main

import (
	"context"
	"fmt"
	"net/http"

	"consumer/app/config"
	"consumer/app/connections"
	"consumer/app/handlers"
)

func main() {
	fmt.Println("Setting up server in 3000")
	http.HandleFunc("/socket", handlers.Upgrade)
	ctx := context.Background()
	go config.Listen(ctx, connections.SendMessage)
	http.ListenAndServe(":3000", nil)
}
