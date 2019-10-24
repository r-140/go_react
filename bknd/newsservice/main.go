package main

import (
	"fmt"

	"dbclient"
	"service"
)

var appName = "newsservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6768")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenDbClient()
	service.DBClient.Seed()
}
