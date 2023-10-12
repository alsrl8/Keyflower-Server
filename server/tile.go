package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
)

func putTile(hub *Hub, tile *game.Tile) {
	data, _ := json.Marshal(tile)
	message := Message{
		Type: enum.NewTile,
		Data: string(data),
	}
	hub.broadcast(message)
}
