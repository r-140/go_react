package main

import (
	"fmt"

	"dbclient"
	"service"
)

var appName = "newsservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("6767")
}

func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenDbClient()
	service.DBClient.Seed()
}
