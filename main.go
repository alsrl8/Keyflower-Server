package main

import (
	"Server/game"
	"Server/server"
)

func init() {
	game.GenerateInitialTiles()
	game.GenerateInitialMeeples()
}

func main() {
	server.Run()
}
