package server

import (
	"Server/enum"
	"Server/game"
	"Server/utils"
	"log"
)

func (hub *Hub) handleMeepleAction(meepleActionData *MeepleActionData) {
	for _, detailMeepleAction := range meepleActionData.DetailMeepleActions {
		parentMeepleID := detailMeepleAction.MeepleID
		game.SetChildrenMeeples(parentMeepleID, detailMeepleAction.ChildrenMeepleIDs)
		log.Printf("Meeple Group: parent(%s): %+v", parentMeepleID, game.ChildrenMeepleMap[detailMeepleAction.MeepleID])
	}

	data := utils.ConvertStructToJsonString(meepleActionData)
	message := Message{
		Type: enum.MeepleAction,
		Data: string(data),
	}
	exclusive := []string{meepleActionData.PlayerID}
	hub.broadcastWithExclusivePlayerID(&message, exclusive)
	hub.nextTurn()
	hub.sendTurnChangeData()
}
