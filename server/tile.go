package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
)

func (hub *Hub) putTile(tile *game.Tile) {
	data, _ := json.Marshal(tile)
	message := Message{
		Type: enum.NewTile,
		Data: string(data),
	}
	hub.broadcast(&message)
}
