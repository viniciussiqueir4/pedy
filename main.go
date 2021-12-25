package main

import (
	"pedy/config"
	"pedy/database"
	"pedy/server"
)

func main() {
	config.Init()
	database.StartDatabase()
	s := server.NewServer()
	s.Run()
}
