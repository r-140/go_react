package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

// Signup ...
func Signup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println("CreateNews: user from body ", user)

	result, error := DBClient.CreateUser(user)

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

// Signin signin request handler
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// var result model.User
	result, error := DBClient.QueryUser(creds.Username)

	if error != nil || !result.ComparePwd(creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// Refresh refresh request handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// (END) The code uptil this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
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
