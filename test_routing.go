package main

import (
	"fmt"
	"log"
	"net/http"
	"terame.com/mschuett/lucy"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	temp := lucy.Diamond()
	//fmt.Println(temp)
	temp.Get("/test", http.HandlerFunc(HelloServer))
	temp.Get("/test1", http.HandlerFunc(HelloServer))
	temp.Get("/test2", http.HandlerFunc(HelloServer))
	temp.Post("/post", http.HandlerFunc(HelloServer))
	http.Handle("/", temp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
