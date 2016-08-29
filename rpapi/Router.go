/*
Configure mux.Router with application routes from Routes.go

Call NewRouter() to get the fully configured Router
Attaches the Logger middleware to the pipeline
*/

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Preferred route prefix
//[HOSTNAME].[com]/[prefix]/[route]
//Allows typical prefix of "V1" or "api"
//eg. "/api"
const prefix string = "/api"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//Creates mux.Router object which implements http.Handler
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
