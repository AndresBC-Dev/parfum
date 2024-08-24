package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndresBC-Dev/parfum/middleware"
	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!, this is the first response.")
}

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/", HandleHello)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
