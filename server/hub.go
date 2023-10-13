package server

import (
	"Server/enum"
	"Server/game"
	"Server/utils"
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
	data := utils.ConvertStructToJsonString(message)
	err := client.WriteMessage(websocket.TextMessage, data)
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
		if utils.IsStructInSlice(client, exclusive) {
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
		if utils.IsStructInSlice(playerID, exclusive) {
			continue
		}
		err := hub.sendMessageToClient(client, message)
		if err != nil {
			delete(hub.clients, client)
		}
	}
}

func (hub *Hub) handleGameReady() {
	hub.assignMeepleToPlayer()
	hub.distributeInitialMeepleToPlayer()

	// TODO 현재는 6개로 고정하고 나중에 플레이어 숫자에 따라 초기 타일 숫자를 조절할 것
	tiles := game.GetInitialSpringTiles(6)
	for _, tile := range tiles {
		hub.putTile(tile)
	}

	hub.currentSeason = enum.Spring
	hub.SendGameReadySignal()
	hub.nextTurn()
	hub.sendTurnChangeData()
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
		utils.ConvertJsonStringToStruct(&serverMessage, msg)

		switch serverMessage.Type {
		case enum.Register:
			hub.handleRegister(ws)
			if len(hub.clients) >= hub.PlayerNum {
				log.Println("All players entered!")
				hub.handleGameReady()
			}
		case enum.MeepleAction:
			data := MeepleActionData{}
			utils.ConvertJsonStringToStruct(&data, []byte(serverMessage.Data))
			log.Printf("Success!! -> %+v", data) // TODO Meeple Action을 핸들링하는 로직 추가
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
