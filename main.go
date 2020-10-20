package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/anz-bank/sysl/pkg/parse"
	"github.com/anz-bank/sysl/pkg/sysl"
	"github.com/sandertv/mcwss"
	"github.com/sandertv/mcwss/mctype"
)

var initpos mctype.Position
var initialized bool
var uniqueIDs []string
var selectednamespaces []string

var agent mcwss.Agent
var namespacesp []mctype.Position

func main() {
	// Create a new server using the default configuration. To use specific configuration, pass a *wss.Config{} in here.
	server := mcwss.NewServer(&mcwss.Config{
		HandlerPattern: "/ws",
		Address:        "0.0.0.0:" + os.Getenv("PORT"),
	})

	p := parse.NewParser()
	module, _ := p.ParseString(`
App:
    _:
		foobar:
			return ok <: string
App2:
    _:
        oasdoiasd:
            return ok <: string
App3:
    _:
        oasdoiasd:
            return ok <: string`)
	server.OnConnection(func(player *mcwss.Player) {
		MOTD(player)
		InitArea(player)
		fmt.Println(player)
		fmt.Println("Here2")
		DoSysl(player, module)
		fmt.Println("Here3")
	})
	server.OnDisconnection(func(player *mcwss.Player) {
		fmt.Println(player)
		// Called when a player disconnects from the server.
	})
	// Run the server. (blocking)
	server.Run()
}

// MOTD will display the title and subtitle
func MOTD(player *mcwss.Player) {
	player.Exec(fmt.Sprintf("title %s title SYSL", player.Name()), nil)
	player.Exec(fmt.Sprintf("title %s subtitle The best specification language in all of the world", player.Name()), nil)
}

func DoSysl(p *mcwss.Player, m *sysl.Module) {
	fmt.Println(m)
	p.Position(func(pos mctype.Position) {
		for _, app := range m.Apps {
			fmt.Println("Here2")
			p.Exec(fmt.Sprintf("summon %s ~ ~ ~ %s %s", "chicken", app.Name, app.Name), nil)
		}
	})
}

func InitArea(p *mcwss.Player) {
	p.Position(func(pos mctype.Position) {
		// Create animal pens
		Fill(p, pos, -20, -2, -20, 20, 15, 20, "air")
		Fill(p, pos, -15, -2, -15, 15, -1, 15, "stone 4")
		Fill(p, pos, -1, -2, -1, 1, -2, 1, "glass")
		Fill(p, pos, 0, -2, 0, 0, -2, 0, "beacon")
		Fill(p, pos, -14, -1, -14, 14, -1, 14, "air")
		Fill(p, pos, -14, -2, -14, -2, -2, -2, "grass")
		Fill(p, pos, -14, -1, -14, -2, -1, -2, "fence")
		Fill(p, pos, -13, -1, -13, -9, -1, -9, "air")
		Fill(p, pos, -13, -1, -7, -9, -1, -3, "air")
		Fill(p, pos, -7, -1, -13, -3, -1, -9, "air")
		Fill(p, pos, -7, -1, -7, -3, -1, -3, "air")

		initpos = pos

		namespacesp = []mctype.Position{
			{X: pos.X - 11, Y: pos.Y + 5, Z: pos.Z - 11},
			{X: pos.X - 11, Y: pos.Y + 5, Z: pos.Z - 5},
			{X: pos.X - 5, Y: pos.Y + 5, Z: pos.Z - 11},
			{X: pos.X - 5, Y: pos.Y + 5, Z: pos.Z - 5},
		}
	})
}

// Fill will fill the playing area with blocktype, coordinates are absolute
func Fill(p *mcwss.Player, pos mctype.Position, x1 int, y1 int, z1 int, x2 int, y2 int, z2 int, blocktype string) {
	p.Exec(fmt.Sprintf("fill %d %d %d %d %d %d %s", int(pos.X)+x1, int(pos.Y)+y1, int(pos.Z)+z1, int(pos.X)+x2, int(pos.Y)+y2, int(pos.Z)+z2, blocktype), nil)
}

// Summonpos will spawn a named entity in a random area close to the position passed - UniqueID check will prevent spawning an entity more than once
func Summonpos(p *mcwss.Player, pos mctype.Position, entity string, name string) {
	if !Contains(uniqueIDs, name) {
		uniqueIDs = append(uniqueIDs, name)
		p.Exec(fmt.Sprintf("summon %s %s %d %d %d", entity, name, int(pos.X-1.5+3*rand.Float64()), int(pos.Y)-5, int(pos.Z-1.5+3*rand.Float64())), nil)
		time.Sleep(100 * time.Millisecond)
	} else {
		fmt.Printf("Entity %s already exists\n", name)
	}
}

// Contains checks for the occurence of a string in an array of strings
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
