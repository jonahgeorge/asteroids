package main

import (
	"github.com/lucasb-eyer/go-colorful"
)

type Vector2f struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Ship struct {
	Position     Vector2f `json:"position"`
	Acceleration Vector2f `json:"acceleration"`
	Velocity     Vector2f `json:"velocity"`
	Heading      float64  `json:"heading"`
	Thrust       float64  `json:"thrust"`
	TurnSpeed    float64  `json:"turn_speed"`
	Color        string   `json:"color"`
}

type Client struct {
	Name string `json:"name"`
	Ship Ship   `json:"ship"`
}

func NewClient() Client {
	var client Client
	client.Name = "Anonymous"
	client.Ship.Position.X = 50
	client.Ship.Position.Y = 50
	client.Ship.Heading = 0
	client.Ship.Thrust = 1
	client.Ship.TurnSpeed = 0.3

	colors := colorful.FastHappyPalette(1)
	client.Ship.Color = colors[0].Hex()

	return client
}
