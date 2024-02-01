package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/Edu4rdoNeves/Address-search-api/infrastructure/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	server := server.New()
	server.Run()
}
