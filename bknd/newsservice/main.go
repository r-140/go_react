package main

import (
	"log"

	"dbclient"
	"os"
	"service"
)

var appName = "newsservice"

func main() {
	log.Printf("Starting %v\n", appName)
	port := os.Getenv("PORT")

	initializeMongoClient()
	service.StartWebServer(port)
}

func initializeMongoClient() {
	dbclient.DBClient = &dbclient.MongoClient{}
	dbclient.DBClient.OpenDbClient()
	// service.DBClient.Seed()
}
