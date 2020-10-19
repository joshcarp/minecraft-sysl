package main

import (
	"fmt"
	"github.com/sandertv/mcwss"
)

func main() {
	// Create a new server using the default configuration. To use specific configuration, pass a *wss.Config{} in here.
	server := mcwss.NewServer(nil)

	server.OnConnection(func(player *mcwss.Player){
		MOTD(player)
		// Called when a player connects to the server.

	})
	server.OnDisconnection(func(player *mcwss.Player){
		// Called when a player disconnects from the server.
	})
	// Run the server. (blocking)
	server.Run()
}
// MOTD will display the title and subtitle
func MOTD(player *mcwss.Player) {
	player.Exec(fmt.Sprintf("title %s title KubeCraftAdmin", player.Name()), nil)
	player.Exec(fmt.Sprintf("title %s subtitle The Adventurer's Admin Tool", player.Name()), nil)
}