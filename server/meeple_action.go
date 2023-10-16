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
	}

	for _, detailMeepleAction := range meepleActionData.DetailMeepleActions {
		tileID := detailMeepleAction.TargetTileID
		if game.TileMap[tileID].BidNum < detailMeepleAction.Number {
			game.TileMap[tileID].BidNum = detailMeepleAction.Number
			game.TileMap[tileID].OwnerID = meepleActionData.PlayerID
			log.Printf("Tile Owner Changed. Tile(%s) BidNum(%d) Owner(%s)", tileID, game.TileMap[tileID].BidNum, game.TileMap[tileID].OwnerID)
		}
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
