package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
	"log"
)

func (hub *Hub) nextTurn() {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	hub.currentTurn++
	if hub.currentTurn > hub.PlayerNum {
		hub.currentTurn = 1
	}
}

func (hub *Hub) sendTurnChangeData() {
	data, _ := json.Marshal(TurnChangeData{Turn: hub.currentTurn})
	message := Message{Type: enum.TurnChange, Data: string(data)}
	hub.broadcast(&message)
}

func (hub *Hub) nextSeason() {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	switch hub.currentSeason {
	case enum.Spring:
		tiles := game.GetRoundTilesBySeason(enum.Summer, 6)
		for _, tile := range tiles {
			hub.putTile(tile)
		}
		hub.currentSeason = enum.Summer
	case enum.Summer:
		hub.currentSeason = enum.Autumn
	case enum.Autumn:
		hub.currentSeason = enum.Winter
	case enum.Winter:
		// TODO 겨울 라운드가 끝나고 점수 계산하는 기능 추가
		break
	}
	hub.playerSkipCnt = 0
	hub.sendSeasonChangeData()
}

func (hub *Hub) sendSeasonChangeData() {
	data, err := json.Marshal(SeasonChangeData{Season: hub.currentSeason})
	if err != nil {
		log.Printf("Failed to parse SeasonChangeData into json string: %+v", err)
	}
	message := Message{Type: enum.SeasonChange, Data: string(data)}
	hub.broadcast(&message)
}
