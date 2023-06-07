package api

import (
	"encoding/json"
	"log"
	"net/http"
	"service/models"

	"github.com/gorilla/mux"
)

type server struct {
	jobs models.JobRepository
}

func New(jobs models.JobRepository) *mux.Router {
	server := newServer(jobs)
	return newRouter(server)
}

func newServer(jobs models.JobRepository) *server {
	return &server{jobs}
}

// healthCheck is a handler function that returns a 200 status code
func (srv *server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Health check!"))
	log.Printf("Health check")
}

// getJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *server) getJob(w http.ResponseWriter, r *http.Request) {
	// create a new job
	payload := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := payload.FromJSON(r.Body)
	if err != nil {
		log.Printf("Unable to unmarshal JSON")
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	job, err := srv.jobs.FindJobByID(payload.JobId)
	// Encode the User object into JSON
	jsonData, err := json.Marshal(job)
	if err != nil {
		log.Printf("Internal server Error")
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL with job
	// Write the JSON response
	w.Write(jsonData)
	log.Printf("Get Job: %v", job)

}

// getJobs is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *server) getJobs(w http.ResponseWriter, r *http.Request) {
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Get Jobs!"))
	log.Printf("Get Jobs")
}

// createJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *server) createJob(w http.ResponseWriter, r *http.Request) {
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Create Job!"))
	log.Printf("Create Job")
}

// updateJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *server) updateJob(w http.ResponseWriter, r *http.Request) {
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Update Job!"))
	log.Printf("Update Job")
}

// deleteJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *server) deleteJob(w http.ResponseWriter, r *http.Request) {
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Delete Job!"))
	log.Printf("Delete Job")
}
