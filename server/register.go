package server

import (
	"Server/enum"
	"Server/utils"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func (hub *Hub) handleRegister(client *websocket.Conn) {
	newID := utils.GenerateRandomID()
	err := hub.sendMessageToClient(client, &Message{
		enum.Register,
		newID,
	})
	if err != nil {
		return
	}
	hub.clients[client] = newID
	hub.NotifyNewClient(newID)
}

// NotifyNewClient Notify all clients there is newly added client.
func (hub *Hub) NotifyNewClient(newClientID string) {
	newPlayerData, _ := json.Marshal(NewPlayerData{NewPlayerID: newClientID, TotalPlayerNum: len(hub.clients)})
	message := Message{Type: enum.NewPlayer, Data: string(newPlayerData)}
	hub.broadcast(&message)
}
