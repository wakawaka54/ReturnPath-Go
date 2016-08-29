/*
Logger Middleware Handler

Takes in an HTTP Requests and starts timing.
After HTTP.Request executes, it logs requests and outputs execution time

Eg.
GET			api/sentences			GetSentences			10ms
*/


package main

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Start the clock!
		start := time.Now()

		inner.ServeHTTP(w, r)

		//Continues after inner handler executes
		fmt.Printf(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
