package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"dbclient"
	"model"

	"github.com/gorilla/mux"
)

// Claims represent claims into jwt
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

var DBClient dbclient.IDbClient

var isHealthy = true

var client = &http.Client{}

func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
}

// GetNewsById handles getnews request
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

// CreateNews ...
func CreateNews(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	decoder := json.NewDecoder(r.Body)
	var news model.News
	err := decoder.Decode(&news)
	if err != nil {
		panic(err)
	}

	fmt.Println("CreateNews: news from body ", news)

	result, error := DBClient.CreateNews(news)

	if error != nil {
		fmt.Println("Some error occured creating news " + error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(result)
	writeJsonResponse(w, http.StatusCreated, data)

}

// CreateComment ...
func CreateComment(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var newsID = mux.Vars(r)["newsID"]

	decoder := json.NewDecoder(r.Body)
	var comment model.Comment
	err := decoder.Decode(&comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("CreateComment: news from body ", comment)

	result, error := DBClient.CreateComment(newsID, comment)

	if error != nil {
		fmt.Println("Some error occured creating news " + error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(result)
	writeJsonResponse(w, http.StatusCreated, data)

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
