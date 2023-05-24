package main

import (
	"log"
	"net/http"
	"service/api"
)

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.Lshortfile)
	log.Printf("STARTING Server")

	// setup router
	router := api.NewRouter()

	go func() {
		log.Println(http.ListenAndServe(":8080", router))
	}()
	select {}
}
