package server

import "os"

func New() Server {
	server := NewServer()

	server.SetPort(os.Getenv("SERVER_PORT"))
	server.SetHealthMessage("Health Response Okay")
	return *server
}
