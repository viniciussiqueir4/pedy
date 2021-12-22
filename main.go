package main

import (
	"pedy/config"
	"pedy/server"
)

func main() {
	config.Init()
	s := server.NewServer()
	s.Run()
}
