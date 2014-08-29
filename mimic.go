package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mimic is simple mock server for no-engineer.
type Mimic struct {
	port   int
	config *Config
	router *mux.Router
}

// NewMimic returns mimic instance.
func NewMimic(p int, c *Config) Mimic {
	r := routing(c.Rules)
	m := Mimic{
		port:   p,
		config: c,
		router: r,
	}

	return m
}

// Start mimic server
func (m *Mimic) Start() {
	http.Handle("/", m.router)
	log.Printf("mimic server started at localhost:%d", m.config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", m.port), nil)
}
