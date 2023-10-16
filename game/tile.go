package game

import (
	"Server/enum"
	"Server/utils"
	"encoding/json"
	"fmt"
	"os"
)

type Tile struct {
	ID        string `json:"tileID"`
	OwnerID   string `json:"ownerID"`
	BidNum    int
	*TileInfo `json:"tileInfo"`
}

type TileInfo struct {
	Name             string `json:"name"`
	enum.Season      `json:"season"`
	IsUpgraded       bool           `json:"isUpgraded"`
	CostToUpgrade    map[string]int `json:"costToUpgrade"`
	BasicTileInfo    DetailTileInfo `json:"basicTileInfo"`
	UpgradedTileInfo DetailTileInfo `json:"upgradedTileInfo"`
}

type DetailTileInfo struct {
	Point  int            `json:"point"`
	Cost   map[string]int `json:"cost"`
	Reward map[string]int `json:"reward"`
}

var TileMap map[string]*Tile

func init() {
	TileMap = make(map[string]*Tile)
}

func GenerateInitialTiles() {
	data, err := os.ReadFile("./game/tileinfo.json")
	if err != nil {
		fmt.Println(err)
	}
	var tileInfos []TileInfo
	err = json.Unmarshal(data, &tileInfos)
	if err != nil {
		fmt.Println(err)
	}

	tileNum := len(tileInfos)
	for i := 0; i < tileNum; i++ {
		tile := Tile{
			ID:       utils.GenerateRandomID(),
			OwnerID:  "",
			TileInfo: &tileInfos[i],
		}
		TileMap[tile.ID] = &tile
	}
}

func getSeasonTileIDs(season enum.Season) (tileIDs []string) {
	for tileID, tile := range TileMap {
		if tile.Season != season {
			continue
		}
		tileIDs = append(tileIDs, tileID)
	}
	return
}

func GetRoundTilesBySeason(season enum.Season, requiredNum int) (tiles []*Tile) {
	springTileIDs := getSeasonTileIDs(season)
	utils.ShuffleSlice(springTileIDs)
	for _, tileID := range springTileIDs {
		tiles = append(tiles, TileMap[tileID])
		if len(tiles) == requiredNum {
			return
		}
	}
	return
}
