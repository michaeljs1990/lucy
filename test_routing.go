package main

import (
	"fmt"
	"log"
	"net/http"
	"terame.com/mschuett/lucy"
)

func HelloServer(w lucy.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	temp := lucy.Diamond()
	temp.Get("/test", http.HandlerFunc(HelloServer))
	http.Handle("/", temp)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
