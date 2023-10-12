package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
)

func assignMeepleToPlayer(hub *Hub) {
	playerIDs := make([]string, 0)
	for _, playerID := range hub.clients {
		playerIDs = append(playerIDs, playerID)
	}
	game.AssignMeepleToPlayer(playerIDs)
}

func getAllMeeple() []string {
	return game.GetAllMeepleID()
}

func distributeInitialMeepleToPlayer(hub *Hub) {
	for client, _ := range hub.clients {
		//meepleIDs := getAllMeepleIDsByPlayerID(playerID)
		meepleIDs := getAllMeeple()

		for _, meepleID := range meepleIDs {
			data, err := json.Marshal(game.Meeple{
				ID:      meepleID,
				OwnerID: game.MeepleMap[meepleID].OwnerID,
				Color:   game.MeepleMap[meepleID].Color,
			})
			if err != nil {
				continue
			}
			message := Message{
				Type: enum.NewMeeple,
				Data: string(data),
			}
			err = hub.sendMessageToClient(client, message)
			if err != nil {
				continue
			}
		}
	}
}

func sendSignalMeepleMovement(hub *Hub, playerID string, moveMeepleData MoveMeepleData) {
	data := OtherPlayerActionData{
		PlayerID: playerID,
		Actions:  make([]PlayerActionData, 0),
	}
	data.Actions = append(data.Actions, PlayerActionData{
		Type: enum.MoveMeeple,
		Data: convertStructToJsonString(moveMeepleData),
	})
	message := Message{
		Type: enum.OtherPlayerAction,
		Data: convertStructToJsonString(data),
	}

	exclusive := make([]string, 0)
	exclusive = append(exclusive, playerID)
	hub.broadcastWithExclusivePlayerID(message, exclusive)
}
