package main

import (
	//"fmt"
	"log"
	"net/http"
	"terame.com/mschuett/lucy"
)

// func HelloServer(r Path) {
// 	fmt.Fprintf(w, "hello")
// }

func testing(s lucy.Params) ([]byte, int) {
	return []byte(s.Get("this")), 503
}

func vans(s lucy.Params) ([]byte, int) {
	return []byte(s.Get("id")), 404
}

func users(s lucy.Params) ([]byte, int) {
	return []byte(s.Get("this")), 200
}

func main() {
	temp := lucy.Diamond()
	temp.Get("/test", lucy.HandlerFunc(testing))
	temp.Get("/vans/:id", lucy.HandlerFunc(vans))
	temp.Put("/users/:epanther", lucy.HandlerFunc(users))
	http.Handle("/", temp)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
