package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routing(rules []Rule) *mux.Router {
	r := mux.NewRouter()

	for _, rule := range rules {
		registerRoute(r, rule)
		log.Printf("registered: %s", rule)
	}

	return r
}

func registerRoute(r *mux.Router, rule Rule) {
	r.HandleFunc(rule.Path, func(res http.ResponseWriter, req *http.Request) {
		log.Println(rule)
		res.Write([]byte(rule.Content))
	}).Methods(rule.Method).Name(rule.String())
}
