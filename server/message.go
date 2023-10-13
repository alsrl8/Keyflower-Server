package server

import (
	"Server/enum"
)

// Message corresponds to the C# class for handling messages from the server.
type Message struct {
	Type enum.ServerMessageType `json:"type"`
	Data string                 `json:"data"`
}

type NewPlayerData struct {
	NewPlayerID    string `json:"newPlayerID"`
	TotalPlayerNum int    `json:"totalPlayerNum"`
}

type GameReadyData struct {
	PlayerID   string `json:"playerID"`
	PlayerTurn int    `json:"playerTurn"`
}

type TurnChangeData struct {
	Turn int `json:"turn"`
}

type SeasonChangeData struct {
	enum.Season `json:"season"`
}
