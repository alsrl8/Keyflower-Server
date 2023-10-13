package server

import (
	"Server/enum"
	"encoding/json"
)

func (hub *Hub) SendGameReadySignal() {
	turn := 0
	for client, clientID := range hub.clients {
		turn += 1
		data, _ := json.Marshal(GameReadyData{
			PlayerID:   clientID,
			PlayerTurn: turn,
		})
		message := Message{Type: enum.GameReady, Data: string(data)}
		_ = hub.sendMessageToClient(client, &message)
	}
}
