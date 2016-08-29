/*
Allows Application Pipeline to be configured per Middleware Model

eg.

EnableCorsMiddleware(next http.Handler) http.Handler {
//Execute this before the next
fmt.Println("Before next request")

//Must call next to execute downstream Handler
//Can short circuit the request if next is not called
next.ServeHTTP(w,r)

//Continues executing the rest of the handler after next is called
fmt.Println("Wow, that was a lot of work")
}

Configuration:

middleware := Middleware{}
pipeline = middleware.
							AddServie(func (next http.Handler) http.Handler). //Add middleware FILO
							AddHandler(func () http.Handler). //Add a final handler with no next
							Build() //Return final http.Handler based Pipeline
*/

package main

import (
	"net/http"
)

type Middleware struct {
	handler  http.Handler
	pipeline [](func(next http.Handler) http.Handler)
}

//Add middleware
func (m *Middleware) AddService(f func(http http.Handler) http.Handler) *Middleware {
	m.pipeline = append(m.pipeline, f)
	return m
}

//Add final handler
func (m *Middleware) UseHandler(r http.Handler) *Middleware {
	m.handler = r
	return m
}

//Create pipeline
func (m *Middleware) Build() http.Handler {
	if len(m.pipeline) == 0 {
		return nil
	}

	handler := m.handler

	for _, pipe := range m.pipeline {
		handler = pipe(handler)
	}

	return handler
}
