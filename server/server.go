package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"pedy/config"
	"pedy/database"
	"pedy/server/routes"
)

type Server struct {
	port           string
	server         *gin.Engine
	DB             *gorm.DB
}

func NewServer() Server {
	return Server{
		port:           config.GetConfig().ServerPort,
		server:         gin.Default(),
		DB:             database.StartDatabase(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server, s.DB)
	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

