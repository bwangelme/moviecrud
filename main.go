package main

import (
	"log"
	"moviedemo/movieview/router"
	"net/http"
)

func main() {
	r := router.Load()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
