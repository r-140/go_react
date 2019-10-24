package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"dbclient"

	"github.com/gorilla/mux"
)

var DBClient dbclient.IDbClient

var isHealthy = true

var client = &http.Client{}

func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
}

func GetNewsById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Read the 'newsId' path parameter from the mux map
	var newsID = mux.Vars(r)["newsID"]

	// Read the news struct BoltDB
	news, err := DBClient.QueryNews(newsID)

	// If err, return a 404
	if err != nil {
		fmt.Println("Some error occured serving " + newsID + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(news)
	writeJsonResponse(w, http.StatusOK, data)
}

// GetAllnews ...
func GetAllnews(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	// Read the news struct BoltDB
	news, err := DBClient.QueryAllNews()

	// If err, return a 404
	if err != nil {
		fmt.Println("Some error occured serving " + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(news)
	writeJsonResponse(w, http.StatusOK, data)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
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
