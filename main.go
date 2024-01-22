package main

import (
	"log"
	"runners-postresql/config"
	"runners-postresql/server"

	_ "github.com/pq"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing Configuration")
	config := config.InitConfig("runners")
	log.Println("Initializing Database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
