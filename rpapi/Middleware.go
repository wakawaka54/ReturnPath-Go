package main

import (
	"net/http"
)

type Middleware struct {
	handler  http.Handler
	pipeline [](func(next http.Handler) http.Handler)
}

func (m *Middleware) AddService(f func(http http.Handler) http.Handler) *Middleware {
	m.pipeline = append(m.pipeline, f)
	return m
}

func (m *Middleware) UseHandler(r http.Handler) *Middleware {
	m.handler = r
	return m
}

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
