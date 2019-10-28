package service

import (
	"graphql"

	"github.com/gorilla/mux"
)

// func NewRouter() *mux.Router {

// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {

// 		router.Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(route.HandlerFunc)

// 	}
// 	return router
// }

// NewRouter creates a mux.Router pointer.
func NewRouter() *mux.Router {

	graphql.InitQL(&graphql.LiveGraphQLResolvers{})

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		// Handler(loadTracing(route.HandlerFunc, route.Name))
	}
	return router
}
