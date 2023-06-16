package main

import (
	"log"
	"net/http"
	"service/api"
	_ "service/docs" // This line is necessary for go-swagger to find your docs!
	"service/repositories"
)

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.Lshortfile)
	log.Printf("STARTING Service")

	// create a new pack using the database connection
	//pack := repositories.Pack{Jobs: &repositories.JobRepository(db)}
	jobs := repositories.NewJobsMemory()
	// setup router
	server := api.NewServer(jobs)
	router := api.NewRouter(server)

	go func() {
		log.Println(http.ListenAndServe(":8080", router))
	}()
	select {}
}
