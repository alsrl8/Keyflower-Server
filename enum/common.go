package enum

type ServerMessageType string

const (
	CommonMessage     ServerMessageType = "CommonMessage"
	Register          ServerMessageType = "Register"
	NewPlayer         ServerMessageType = "NewPlayer"
	GameReady         ServerMessageType = "GameReady"
	NewMeeple         ServerMessageType = "NewMeeple"
	NewTile           ServerMessageType = "NewTile"
	TurnChange        ServerMessageType = "TurnChange"
	SeasonChange      ServerMessageType = "SeasonChange"
	EndPlayerAction   ServerMessageType = "EndPlayerAction"
	OtherPlayerAction ServerMessageType = "OtherPlayerAction"
)

type PlayerActionType string

const (
	MoveMeeple PlayerActionType = "MoveMeeple"
)

type MeepleColor string

const (
	Red    MeepleColor = "Red"
	Blue   MeepleColor = "Blue"
	Yellow MeepleColor = "Yellow"
	Green  MeepleColor = "Green"
	Purple MeepleColor = "Purple"
)

type Season string

const (
	Spring Season = "Spring"
	Summer Season = "Summer"
	Autumn Season = "Autumn"
	Winter Season = "Winter"
)
