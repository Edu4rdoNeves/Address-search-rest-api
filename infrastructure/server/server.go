package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port          string
	healthMessage string
	serverEngine  *gin.Engine
}

func NewServer() *Server {
	return &Server{
		port:          os.Getenv("SERVER_PORT"),
		healthMessage: "",
		serverEngine:  gin.Default(),
	}
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) SetHealthMessage(healthMessage string) {
	s.healthMessage = healthMessage
}

func (s *Server) helathRoute(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, s.healthMessage)
}

func (s *Server) SetServer(server *gin.Engine) {
	s.serverEngine = server
}

func (s *Server) Run() {
	router := ConfigRouters(s.serverEngine)

	s.serverEngine.GET("/health", s.helathRoute)

	var adress string = ":" + s.port

	s.serverEngine.Run(adress)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
