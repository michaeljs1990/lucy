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

func testing(s *lucy.Params) []byte {
	return []byte(s.Get("this"))
}

func main() {
	temp := lucy.Diamond()
	temp.Get("/test", lucy.HandlerFunc(testing))
	http.Handle("/", temp)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
