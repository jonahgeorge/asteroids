package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/googollee/go-socket.io"
)

var players []Player

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Ship struct {
	Position     Point   `json:"position"`
	Acceleration Point   `json:"acceleration"`
	Velocity     Point   `json:"velocity"`
	Heading      float64 `json:"heading"`
	Thrust       float64 `json:"thrust"`
	TurnSpeed    float64 `json:"turn_speed"`
}

type Player struct {
	Name string `json:"name"`
	Ship Ship   `json:"ship"`
}

func UpdatePlayers(socket *socketio.SocketIOServer) {
	for _ = range time.Tick(30 * time.Millisecond) {
		for key, _ := range players {
			// Update location
			players[key].Ship.Position.X += players[key].Ship.Velocity.X
			players[key].Ship.Position.Y += players[key].Ship.Velocity.Y
		}
		// fmt.Printf("%+v\n", players)
		bytes, _ := json.Marshal(players)
		socket.Broadcast("update", string(bytes))
	}
}

// Throw some error handling in
func FindPlayerById(id string) int {
	for k, v := range players {
		if v.Name == id {
			return k
		}
	}
	return -1
}

func move(ns *socketio.NameSpace, command string) {
	index := FindPlayerById(ns.Id())

	switch command {
	case "up":
		players[index].Ship.Acceleration.X = players[index].Ship.Thrust * math.Sin(players[index].Ship.Heading)
		players[index].Ship.Acceleration.Y = players[index].Ship.Thrust * math.Cos(players[index].Ship.Heading)
		players[index].Ship.Velocity.X += players[index].Ship.Acceleration.X
		players[index].Ship.Velocity.Y += players[index].Ship.Acceleration.Y

	case "down":
		// nada

	case "left":
		players[index].Ship.Heading += players[index].Ship.TurnSpeed

	case "right":
		players[index].Ship.Heading -= players[index].Ship.TurnSpeed
	}
}

func main() {
	config := &socketio.Config{}
	config.HeartbeatTimeout = 2
	config.ClosingTimeout = 4

	socket := socketio.NewSocketIOServer(config)

	// Handler for new connections, also adds socket.io event handlers
	socket.On("connect", func(ns *socketio.NameSpace) {
		fmt.Println("user " + ns.Id() + " connected.")

		var player Player
		player.Name = ns.Id()
		player.Ship.Position.X = 50
		player.Ship.Position.Y = 50
		player.Ship.Heading = 0
		player.Ship.Thrust = 1
		player.Ship.TurnSpeed = 0.3

		players = append(players, player)
		// fmt.Printf("%+v\n", players)
	})

	socket.On("disconnect", func(ns *socketio.NameSpace) {
		fmt.Println("user " + ns.Id() + " disconnected.")
		index := FindPlayerById(ns.Id())
		players = append(players[:index], players[index+1:]...)
	})

	socket.On("move", move)

	go UpdatePlayers(socket)

	//this will serve a http static file server
	socket.Handle("/", http.FileServer(http.Dir("./public/")))

	//startup the server
	log.Fatal(http.ListenAndServe(":3000", socket))
}
