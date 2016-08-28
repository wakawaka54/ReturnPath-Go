/*
Entry point of application
Sets up Server, IndexRoute and Static file serving
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var ApiAddress string
var fileHttp = http.NewServeMux()

func main() {

	fmt.Println("Reading config files")

	config := BuildConfig()
	ApiAddress = config.ApiAddress

	fmt.Println("Spinning up the server")

	fs := http.FileServer(http.Dir(HomePath + "/static/"))

	pipeline := Logger(StaticFiles(Index()))

	http.Handle("/", pipeline)
	fileHttp.Handle("/", http.StripPrefix("/~/", fs))

	fmt.Printf("Server running on %s:%s\n", config.Hostname, config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Hostname, config.Port), nil))
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		model := IndexModel{
			ApiAddress: ApiAddress,
		}
		t, _ := template.ParseFiles(HomePath + "index.html")
		t.Execute(w, model)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Printf(
			"%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func StaticFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "~") {
			fmt.Println(r.URL.Path)
			fileHttp.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
