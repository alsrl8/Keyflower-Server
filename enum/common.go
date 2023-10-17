package enum

type ServerMessageType string

const (
	CommonMessage ServerMessageType = "CommonMessage"
	Register                        = "Register"
	NewPlayer                       = "NewPlayer"
	GameReady                       = "GameReady"
	NewMeeple                       = "NewMeeple"
	NewTile                         = "NewTile"
	TurnChange                      = "TurnChange"
	SeasonChange                    = "SeasonChange"
	MeepleAction                    = "MeepleAction"
	Chat                            = "Chat"
)

type MeepleColor string

const (
	Red    MeepleColor = "Red"
	Blue               = "Blue"
	Yellow             = "Yellow"
	Green              = "Green"
	Purple             = "Purple"
)

type Season string

const (
	Spring Season = "Spring"
	Summer        = "Summer"
	Autumn        = "Autumn"
	Winter        = "Winter"
)

type MeepleActionType int

const (
	Bid MeepleActionType = iota
	Play
)
