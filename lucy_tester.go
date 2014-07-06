package main

import (
	"log"
	"net/http"
	"terame.com/mschuett/lucy"
)

func main() {
	t := lucy.Kickstart()

	t.Get("/test/:id", func(e *lucy.Service) {
		e.W.Write([]byte(e.Param.Get("id")))
	})

	http.Handle("/", t)

	//Start Server and listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
