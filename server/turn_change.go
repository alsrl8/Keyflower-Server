package server

import (
	"Server/enum"
	"encoding/json"
	"log"
)

func (hub *Hub) sendTurnChangeData() {
	data, _ := json.Marshal(TurnChangeData{Turn: hub.currentTurn})
	message := Message{Type: enum.TurnChange, Data: string(data)}
	hub.broadcast(&message)
}

func sendSeasonChangeData(hub *Hub) {
	data, err := json.Marshal(SeasonChangeData{Season: hub.currentSeason})
	if err != nil {
		log.Printf("Failed to parse SeasonChangeData into json string: %+v", err)
	}
	message := Message{Type: enum.SeasonChange, Data: string(data)}
	hub.broadcast(&message)
}
