package main

import (
	"Server/game"
	"Server/server"
)

func init() {
	game.GenerateInitialTiles()
}

func main() {
	server.Run()
}
