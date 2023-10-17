package server

import (
	"Server/enum"
	"Server/utils"
)

func (hub *Hub) handleChat(chatData *ChatData) {
	data := utils.ConvertStructToJsonString(chatData)
	message := Message{
		Type: enum.Chat,
		Data: string(data),
	}
	exclusive := []string{chatData.PlayerID}
	hub.broadcastWithExclusivePlayerID(&message, exclusive)
}
