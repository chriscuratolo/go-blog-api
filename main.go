package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//router.GET("/", index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
func index() {
	http.Get("http://db.com/")
}
