package api

import (
	"encoding/json"
	"log"
	"net/http"
	"service/models"

	"github.com/gorilla/mux"
)

type Server struct {
	jobs models.JobRepository
}

func NewServer(jobs models.JobRepository) *Server {
	return &Server{jobs}
}

// healthCheck is a handler function that returns a 200 status code
func (srv *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Health check!"))
	log.Printf("Health check")
}

// getJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *Server) getJob(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get Job")
	// get the job id from the URL
	vars := mux.Vars(r)
	// get the job by id
	id, ok := vars["id"]
	if !ok {
		log.Printf("Job ID not found")
		http.Error(w, "Job ID not found", http.StatusBadRequest)
		return
	}
	job, err := srv.jobs.FindJobByID(r.Context(), id)
	// Encode the User object into JSON
	jsonData, err := json.Marshal(job)
	if err != nil {
		log.Printf("Internal Server Error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
func (srv *Server) getJobs(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get Jobs")
	// get all jobs
	jobs, err := srv.jobs.FindJobs(r.Context())
	log.Println(jobs)
	// Encode the User object into JSON
	jsonData, err := json.Marshal(jobs)
	if err != nil {
		log.Printf("Internal Server Error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL with job
	// Write the JSON response
	w.Write(jsonData)
	log.Printf("Get Jobs: %v", jobs)
}

// createJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *Server) createJob(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create Job")
	// create a new job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		log.Printf("Unable to unmarshal JSON")
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	// create a new job
	err = srv.jobs.CreateJob(r.Context(), &job)
	if err != nil {
		log.Printf("Unable to create job")
		http.Error(w, "Unable to create job", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Create Job!"))
	log.Printf("Create Job")
}

// updateJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *Server) updateJob(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update Job")
	// update a job
	job := models.JobRequest{}
	// unmarshal the request body into the job struct
	err := job.FromJSON(r.Body)
	if err != nil {
		log.Printf("Unable to unmarshal JSON")
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	// update job
	err = srv.jobs.UpdateJob(r.Context(), &job)
	if err != nil {
		log.Printf("Unable to update job")
		http.Error(w, "Unable to update job", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Update Job!"))
	log.Printf("Update Job")
}

// deleteJob is a handler function that returns a 200 status code and unmarshales a job from the request body
func (srv *Server) deleteJob(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete Job")
	// delete a job
	// get the job id from the URL
	vars := mux.Vars(r)
	// get the job by id
	id, ok := vars["id"]
	if !ok {
		log.Printf("Job ID not found")
		http.Error(w, "Job ID not found", http.StatusBadRequest)
		return
	}
	// delete job
	err := srv.jobs.DeleteJob(r.Context(), id)
	if err != nil {
		log.Printf("Unable to delete job")
		http.Error(w, "Unable to delete job", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Delete Job!"))
	log.Printf("Delete Job")

}
