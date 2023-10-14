package server

import (
	"Server/enum"
	"encoding/json"
	"log"
)

func (hub *Hub) nextTurn() {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	hub.currentTurn++
	if hub.currentTurn > hub.PlayerNum {
		hub.currentTurn = 1
		hub.playerSkipCnt = 0
	}
}

func (hub *Hub) sendTurnChangeData() {
	data, _ := json.Marshal(TurnChangeData{Turn: hub.currentTurn})
	message := Message{Type: enum.TurnChange, Data: string(data)}
	hub.broadcast(&message)
}

func (hub *Hub) sendSeasonChangeData() {
	data, err := json.Marshal(SeasonChangeData{Season: hub.currentSeason})
	if err != nil {
		log.Printf("Failed to parse SeasonChangeData into json string: %+v", err)
	}
	message := Message{Type: enum.SeasonChange, Data: string(data)}
	hub.broadcast(&message)
}
