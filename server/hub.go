package server

import (
	"Server/enum"
	"Server/game"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var webSocketUpgrade = websocket.Upgrader{}

type Hub struct {
	PlayerNum     int
	currentTurn   int
	currentSeason enum.Season
	playerSkipCnt int
	clients       map[*websocket.Conn]string
	mu            sync.Mutex
}

func (hub *Hub) sendMessageToClient(client *websocket.Conn, message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to deserialization server message: %+v", err)
		return err
	}
	err = client.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("Failed to write message to client[%+v]: %+v", client, err)
		return err
	}
	log.Printf("Send message to player(%+v): %+v", hub.clients[client], message)
	return nil
}

func (hub *Hub) broadcast(message Message) {
	for client := range hub.clients {
		err := hub.sendMessageToClient(client, message)
		if err != nil {
			delete(hub.clients, client)
		}
	}
}

func (hub *Hub) broadcastWithExclusiveConn(message Message, exclusive []*websocket.Conn) {
	for client := range hub.clients {
		if isStructInSlice(client, exclusive) {
			continue
		}
		err := hub.sendMessageToClient(client, message)
		if err != nil {
			delete(hub.clients, client)
		}
	}
}

func (hub *Hub) broadcastWithExclusivePlayerID(message Message, exclusive []string) {
	for client, playerID := range hub.clients {
		if isStructInSlice(playerID, exclusive) {
			continue
		}
		err := hub.sendMessageToClient(client, message)
		if err != nil {
			delete(hub.clients, client)
		}
	}
}

func (hub *Hub) handleGameReady() {
	assignMeepleToPlayer(hub)
	distributeInitialMeepleToPlayer(hub)

	// TODO 현재는 6개로 고정하고 나중에 플레이어 숫자에 따라 초기 타일 숫자를 조절할 것
	tiles := game.GetInitialSpringTiles(6)
	for _, tile := range tiles {
		putTile(hub, tile)
	}

	hub.currentSeason = enum.Spring
	SendGameReadySignal(hub)
	hub.nextTurn()
	sendTurnChangeData(hub)
}

func (hub *Hub) handlePlayerAction(playerID string, endAction PlayerActionData) {
	switch endAction.Type {
	case enum.MoveMeeple:
		moveMeepleData := MoveMeepleData{}
		err := json.Unmarshal([]byte(endAction.Data), &moveMeepleData)
		if err != nil {
			log.Printf("Failed to deserialization move meeple data: %+v", err)
		}
		log.Printf("player(%s): Move Meeple(%s) to Tile(%s)", playerID, moveMeepleData.MeepleID, moveMeepleData.TileID)
		sendSignalMeepleMovement(hub, playerID, moveMeepleData)
	}
}

func (hub *Hub) nextTurn() {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	hub.currentTurn++
	if hub.currentTurn > hub.PlayerNum {
		hub.currentTurn = 1
		hub.playerSkipCnt = 0
	}
}

func (hub *Hub) nextSeason() {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	switch hub.currentSeason {
	case enum.Spring:
		hub.currentSeason = enum.Summer
	case enum.Summer:
		hub.currentSeason = enum.Autumn
	case enum.Autumn:
		hub.currentSeason = enum.Winter
	case enum.Winter:
		// TODO 겨울 라운드가 끝나고 점수 계산하는 기능 추가
		break
	}
	sendSeasonChangeData(hub)
}

func (hub *Hub) run(ws *websocket.Conn) {
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Printf("Failed to close connection: %+v", ws)
		}
	}(ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message: ", err)
			delete(hub.clients, ws)
			break
		}

		log.Println("Message from client: ", string(msg))

		var serverMessage Message
		err = json.Unmarshal(msg, &serverMessage)
		if err != nil {
			log.Printf("Failed to deserialize server message: %+v", err)
			continue
		}

		switch serverMessage.Type {
		case enum.Register:
			hub.handleRegister(ws)
			if len(hub.clients) >= hub.PlayerNum {
				log.Println("All players entered!")
				hub.handleGameReady()
			}
		case enum.EndPlayerAction:
			var endPlayerActions EndPlayerActionData
			data := serverMessage.Data
			_ = json.Unmarshal([]byte(data), &endPlayerActions)
			if len(endPlayerActions.Actions) == 0 {
				hub.playerSkipCnt++
				if hub.playerSkipCnt == hub.PlayerNum {
					hub.playerSkipCnt = 0
					hub.nextSeason()
				}
			} else {
				hub.playerSkipCnt = 0
			}
			for _, playerActionData := range endPlayerActions.Actions {
				hub.handlePlayerAction(endPlayerActions.PlayerID, playerActionData)
			}
			hub.nextTurn()
			sendTurnChangeData(hub)
		}
	}
}

func Run() {

	hub := Hub{
		PlayerNum: 1,
		clients:   make(map[*websocket.Conn]string),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := webSocketUpgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Failed to upgrade request: ", err)
			return
		}
		go hub.run(ws)
	})

	log.Printf("Server is listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
