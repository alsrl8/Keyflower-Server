package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
)

func (hub *Hub) assignMeepleToPlayer() {
	playerIDs := make([]string, 0)
	for _, playerID := range hub.clients {
		playerIDs = append(playerIDs, playerID)
	}
	game.AssignMeepleToPlayer(playerIDs)
}

func getAllMeeple() []string {
	return game.GetAllMeepleID()
}

func (hub *Hub) distributeInitialMeepleToPlayer() {
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
