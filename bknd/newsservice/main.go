package main

import (
	"fmt"

	"dbclient"
	"os"
	"service"
)

var appName = "newsservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	port := os.Getenv("PORT")

	initializeMongoClient()
	service.StartWebServer(port)
}

func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenDbClient()
	// service.DBClient.Seed()
}
