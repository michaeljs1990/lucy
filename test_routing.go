package main

import (
	"log"
	"net/http"
	"terame.com/mschuett/lucy"
)

func testing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(lucy.Param.Get("id")))
}

func vans(w http.ResponseWriter, r *http.Request) {
	//Params.Get("id")
}

func users(w http.ResponseWriter, r *http.Request) {
	//Params.Get("this")
}

func main() {
	temp := lucy.Diamond()
	temp.Get("/test/:id", lucy.HandlerFunc(testing))
	temp.Get("/vans/:id", lucy.HandlerFunc(vans))
	temp.Put("/users/:epanther", lucy.HandlerFunc(users))
	http.Handle("/", temp)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
