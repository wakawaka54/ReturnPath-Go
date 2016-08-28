package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var Pipeline http.Handler

func main() {
	fmt.Println("Prepopulating the dataset")
	PrepopulateDataset()

	fmt.Println("Reading configuration file")
	config := BuildConfig()
	BoringWords = config.BoringWords

	Pipeline = SetupPipeline()

	fmt.Printf("Server started on: %s:%s\n", config.Hostname, config.Port)
	log.Fatal(http.ListenAndServe(config.Hostname+":"+config.Port, Pipeline))
}

func SetupPipeline() http.Handler {
	fmt.Println("Setting up routes")
	r := NewRouter()

	//Create pipeline with EnableCors middleware
	//HttpRequest -> EnableCors -> RestfulHeaders -> AllowOptionsRequest -> r
	m := Middleware{}
	return m.
		UseHandler(r).
		AddService(AllowOptionsRequest).
		AddService(RestfulHeaders).
		AddService(EnableCors).
		Build()
}

/*
Middleware

EnableCors -> Sets correct headers to allow cross origin requests
*/
var CorsHeaders []string

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Expose-Headers", strings.Join(CorsHeaders, ","))

		next.ServeHTTP(w, r)
	})
}

func RestfulHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "text/plain")
		next.ServeHTTP(w, r)
	})
}

func AllowOptionsRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}
