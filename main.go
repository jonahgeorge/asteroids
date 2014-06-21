package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/googollee/go-socket.io"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	FPS = 30 * time.Millisecond
)

var (
	clients []Client
)

func main() {
	socket := socketio.NewSocketIOServer(&socketio.Config{
		HeartbeatTimeout: 2,
		ClosingTimeout:   4,
	})

	// socket.io event handlers
	socket.On("connect", handleConnect)
	socket.On("move", handleMove)
	socket.On("disconnect", handleDisconnect)

	// serve a static file server at root url
	socket.Handle("/", http.FileServer(http.Dir("./public/")))

	// game loop
	go UpdatePlayers(socket)

	// listen on port 3000 and send to socket server
	err := http.ListenAndServe(":3000", socket)
	if err != nil {
		panic(err.Error)
	}
}

func handleConnect(ns *socketio.NameSpace) {
	fmt.Println("user " + ns.Id() + " connected.")

	var player Client
	player.Name = ns.Id()
	player.Ship.Position.X = 50
	player.Ship.Position.Y = 50
	player.Ship.Heading = 0
	player.Ship.Thrust = 1
	player.Ship.TurnSpeed = 0.3

	player.Ship.Color = colorful.WarmColor().Hex()

	clients = append(clients, player)
}

func handleMove(ns *socketio.NameSpace, command string) {
	index := FindClientById(ns.Id())

	switch command {
	case "up":
		clients[index].Ship.Acceleration.X = clients[index].Ship.Thrust * math.Sin(clients[index].Ship.Heading)
		clients[index].Ship.Acceleration.Y = clients[index].Ship.Thrust * math.Cos(clients[index].Ship.Heading)
		clients[index].Ship.Velocity.X += clients[index].Ship.Acceleration.X
		clients[index].Ship.Velocity.Y += clients[index].Ship.Acceleration.Y

	case "down":
		// nada

	case "left":
		clients[index].Ship.Heading += clients[index].Ship.TurnSpeed

	case "right":
		clients[index].Ship.Heading -= clients[index].Ship.TurnSpeed
	}
}

func handleDisconnect(ns *socketio.NameSpace) {
	fmt.Println("user " + ns.Id() + " disconnected.")
	index := FindClientById(ns.Id())
	clients = append(clients[:index], clients[index+1:]...)
}

// Iterate over clients and update ship locations
func UpdatePlayers(socket *socketio.SocketIOServer) {
	for _ = range time.Tick(FPS) {
		for key, _ := range clients {
			// Update location
			clients[key].Ship.Position.X += clients[key].Ship.Velocity.X
			clients[key].Ship.Position.Y += clients[key].Ship.Velocity.Y
		}
		bytes, _ := json.Marshal(clients)
		socket.Broadcast("update", string(bytes))
	}
}

// Retrieve index of player id
// [todo] - Throw some error handling in
func FindClientById(id string) int {
	for k, v := range clients {
		if v.Name == id {
			return k
		}
	}
	return -1
}
