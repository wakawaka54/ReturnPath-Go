/*
Main application sequence is as follows:

1. Prepopulate datasets
2. Read config and set BoringWords (TODO: Move having to set BoringWords functionality somewhere else, maybe make a ConfigInterface for every
	 structure that needs to use the Config object. Like IOptions in .Net Core)
3. Setup Application Pipeline (Middleware -> Mux.Router Handlers)
4. Start Application Server
*/

package main

import (
	"fmt"
	"log"
	"net/http"
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

//Setup Application Pipeline - Middleware http.Handlers -> Handler http.Handler
func SetupPipeline() http.Handler {

	//Setup Mux Router in Router.go utilizing Routes.Go
	fmt.Println("Setting up routes")
	r := NewRouter()

	//Create pipeline with EnableCors middleware
	//HttpRequest -> EnableCors -> RestfulHeaders -> AllowOptionsRequest -> r
	//First In - Last Out
	//Last In - First Out
	m := Middleware{}
	return m.
		UseHandler(r).
		AddService(AllowOptionsRequest).
		AddService(RestfulHeaders).
		AddService(EnableCors).
		Build()
}


//*********** Middleware

//EnableCors -> Sets correct headers to allow cross origin requests
var CorsHeaders []string

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}

//RestfulHeaders -> Sets headers to meet RESTful Header criteria
func RestfulHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "text/plain")
		next.ServeHTTP(w, r)
	})
}

//AllowOptionsRequest -> SHORT CIRCUITS the application in the case of an options Request
//												This allows an AJAX request to get the Cors Configuration
func AllowOptionsRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}
