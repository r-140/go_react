package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"dbclient"
	"model"

	"github.com/gorilla/mux"
)

var client = &http.Client{}

func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
}

// GetNewsById handles getnews request
func GetNewsById(w http.ResponseWriter, r *http.Request) {

	// Read the 'newsId' path parameter from the mux map
	var newsID = mux.Vars(r)["newsID"]

	log.Println("GetNewsById DbClient ", dbclient.DBClient)

	// Read the news struct Mongodb
	news, err := dbclient.DBClient.QueryNews(newsID)

	// If err, return a 404
	if err != nil {
		log.Println("Some error occured serving " + newsID + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(news)
	writeJsonResponse(w, http.StatusOK, data)
}

// CreateNews ...
func CreateNews(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var news model.News
	err := decoder.Decode(&news)
	if err != nil {
		panic(err)
	}

	log.Println("CreateNews: news from body ", news)

	result, error := dbclient.DBClient.CreateNews(news)

	if error != nil {
		log.Println("Some error occured creating news " + error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(result)
	writeJsonResponse(w, http.StatusCreated, data)

}

// CreateComment ...
func CreateComment(w http.ResponseWriter, r *http.Request) {

	var newsID = mux.Vars(r)["newsID"]

	decoder := json.NewDecoder(r.Body)
	var comment model.Comment
	err := decoder.Decode(&comment)
	if err != nil {
		panic(err)
	}

	log.Println("CreateComment: news from body ", comment)

	result, error := dbclient.DBClient.CreateComment(newsID, comment)

	if error != nil {
		log.Println("Some error occured creating news " + error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(result)
	writeJsonResponse(w, http.StatusCreated, data)

}

// GetAllnews ...
func GetAllnews(w http.ResponseWriter, r *http.Request) {

	news, err := dbclient.DBClient.QueryAllNews()

	log.Println("GetAllnews DbClient ", dbclient.DBClient)

	// If err, return a 404
	if err != nil {
		log.Println("Some error occured serving " + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(news)
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
