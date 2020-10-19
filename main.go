package main

import (
	"fmt"
	"github.com/sandertv/mcwss"
	"os"
)

func main() {
	// Create a new server using the default configuration. To use specific configuration, pass a *wss.Config{} in here.
	server := mcwss.NewServer(&mcwss.Config{
		HandlerPattern: "/ws",
		Address:        "0.0.0.0:"+os.Getenv("PORT"),
	})

	server.OnConnection(func(player *mcwss.Player){
		fmt.Println(player)
	})
	server.OnDisconnection(func(player *mcwss.Player){
		fmt.Println(player)
		// Called when a player disconnects from the server.
	})
	// Run the server. (blocking)
	server.Run()
}