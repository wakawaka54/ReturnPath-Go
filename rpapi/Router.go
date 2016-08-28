package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const prefix string = "/api"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler

		//Use Logger Middleware to log requests in and out
		handler = Logger(route.HandlerFunc, route.Name)
		r.
			Methods(route.Method).
			Path(prefix + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return r
}
