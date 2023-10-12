package game

import (
	"Server/enum"
	"Server/utils"
)

type Meeple struct {
	ID      string           `json:"meepleID"`
	OwnerID string           `json:"ownerID"`
	Color   enum.MeepleColor `json:"color"`
}

var MeepleMap map[string]*Meeple

func init() {
	MeepleMap = make(map[string]*Meeple)
}

func getInitialMeepleNumbers() (meeples map[enum.MeepleColor]int) {
	meeples = make(map[enum.MeepleColor]int)
	meeples[enum.Red] = 40
	meeples[enum.Blue] = 40
	meeples[enum.Yellow] = 40
	meeples[enum.Green] = 20
	meeples[enum.Purple] = 1
	return
}

func GenerateInitialMeeples() {
	meeples := getInitialMeepleNumbers()
	for meepleColor, meepleNumber := range meeples {
		for i := 0; i < meepleNumber; i++ {
			newMeeple := &Meeple{
				ID:      utils.GenerateRandomID(),
				OwnerID: "",
				Color:   meepleColor,
			}
			MeepleMap[newMeeple.ID] = newMeeple
		}
	}
	return
}

func getRedBlueYellowMeepleIDs() (meepleIDs []string) {
	for meepleID, meeple := range MeepleMap {
		if meeple.Color != enum.Red && meeple.Color != enum.Blue && meeple.Color != enum.Yellow {
			continue
		}
		meepleIDs = append(meepleIDs, meepleID)
	}
	return
}

func AssignMeepleToPlayer(playerIDs []string) {
	meepleIDs := getRedBlueYellowMeepleIDs()
	utils.ShuffleSlice(meepleIDs)
	for playerIndex, playerID := range playerIDs {
		for num := 0; num < 8; num++ {
			meepleID := meepleIDs[playerIndex*8+num]
			MeepleMap[meepleID].OwnerID = playerID
		}
	}
}

func GetAllMeepleID() []string {
	meepleIDs := make([]string, 0)
	for meepleID, _ := range MeepleMap {
		meepleIDs = append(meepleIDs, meepleID)
	}
	return meepleIDs
}

func GetAllMeepleIDsByPlayerID(playerID string) []string {
	meepleIDs := make([]string, 0)
	for meepleID, meeple := range MeepleMap {
		if meeple.OwnerID == playerID {
			meepleIDs = append(meepleIDs, meepleID)
		}
	}
	return meepleIDs
}
