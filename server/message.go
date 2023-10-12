package server

import (
	"Server/enum"
)

// Message corresponds to the C# class for handling messages from the server.
type Message struct {
	Type enum.ServerMessageType `json:"type"`
	Data string                 `json:"data"`
}

type MoveMeepleData struct {
	MeepleID string `json:"meepleID"`
	TileID   string `json:"tileID"`
}

type SetTileBidNumData struct {
	TileID   string `json:"tileID"`
	PlayerID string `json:"playerID"`
	BidNum   int    `json:"bidNum"`
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

type EndPlayerActionData struct {
	PlayerID string             `json:"playerID"`
	Actions  []PlayerActionData `json:"actions"`
}

type OtherPlayerActionData struct {
	PlayerID string             `json:"playerID"`
	Actions  []PlayerActionData `json:"actions"`
}

type PlayerActionData struct {
	Type enum.PlayerActionType `json:"type"`
	Data string                `json:"data"`
}

type SeasonChangeData struct {
	enum.Season `json:"season"`
}
