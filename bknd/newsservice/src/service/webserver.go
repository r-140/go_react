package service

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func StartWebServer(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)(r))

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
