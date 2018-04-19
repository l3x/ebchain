package main

import (
	"net/http"
)

func newRouter() *http.ServeMux {
	router := http.NewServeMux()

	// routes
	router.Handle("/", index())
	router.Handle("/healthz", healthz())

	return router
}