package api

import (
	"log"
	"net/http"
	"service/models"
)

// healthCheck is a handler function that returns a 200 status code
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Health check!"))
	log.Printf("Health check")
}

// getJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func getJob(w http.ResponseWriter, r *http.Request) {
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Get Job!"))
	log.Printf("Get Job")
}

// getJobs is a handler function that returns a 200 status code and unmarshales a job from the request body
func getJobs(w http.ResponseWriter, r *http.Request) {
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
func createJob(w http.ResponseWriter, r *http.Request) {
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
func updateJob(w http.ResponseWriter, r *http.Request) {
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
func deleteJob(w http.ResponseWriter, r *http.Request) {
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
