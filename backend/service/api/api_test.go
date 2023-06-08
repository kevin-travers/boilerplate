package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service/models"
	"service/repositories"
	"testing"
)

// TestHealthCheck is a unit test for the healthCheck handler function
func TestHealthCheck(t *testing.T) {
	// Prepare a new HTTP GET request to /health_check endpoint
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.healthCheck)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the response body matches the expected message
	expectedResponse := "Health check!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}

// TestGetJob is a unit test for the getJob handler function
func TestGetJob(t *testing.T) {
	// Create a sample job
	job := models.JobRequest{
		JobId:          "1",
		JobName:        "Test",
		JobType:        "jobService",
		JobStatus:      "Open",
		JobPriority:    "High",
		JobDescription: "Job description goes here",
		JobCreatedDate: "2023-04-21",
	}
	// convert job to byte array
	jobJSON, err := json.Marshal(job)
	bodyReader := bytes.NewReader(jobJSON)
	if err != nil {
		t.Errorf("failed to marshal job: %v", err)
	}
	// Prepare a GET request
	req, err := http.NewRequest("GET", "/job", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// setup context
	ctx := context.Background()
	// Add the sample job to the job repository mock
	jobRepo.CreateJob(ctx, &job)

	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.getJob)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response body into a JobRequest instance
	var responseJob models.JobRequest
	err = json.Unmarshal(rr.Body.Bytes(), &responseJob)

	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	// Validate the response job against the expected job

	if responseJob != job {
		t.Errorf("handler returned unexpected job: got %+v, want %+v", responseJob, job)
	}

}

// TestGetJobs is a unit test for the getJobs handler function
func TestGetJobs(t *testing.T) {
	// Create sample jobs
	jobs := []models.JobRequest{
		{
			JobId:          "1",
			JobName:        "Test",
			JobType:        "jobService",
			JobStatus:      "Open",
			JobPriority:    "High",
			JobDescription: "Job description goes here",
			JobCreatedDate: "2023-04-21",
		},
		{
			JobId:          "2",
			JobName:        "Test2",
			JobType:        "jobService",
			JobStatus:      "Open",
			JobPriority:    "High",
			JobDescription: "Job description goes here",
			JobCreatedDate: "2023-04-22",
		},
	}
	// Prepare a GET request
	req, err := http.NewRequest("GET", "/jobs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// setup context
	ctx := context.Background()
	// Add the sample job to the job repository mock
	for index := range jobs {
		jobRepo.CreateJob(ctx, &jobs[index])
	}
	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.getJobs)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response body into a slice of JobRequest instances
	var responseJobs []models.JobRequest
	err = json.Unmarshal(rr.Body.Bytes(), &responseJobs)
	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	// Validate the response jobs against the expected jobs

	// Check if the number of response jobs matches the expected number of jobs
	if len(responseJobs) != len(jobs) {
		t.Errorf("handler returned unexpected number of jobs: got %d, want %d", len(responseJobs), len(jobs))
	}

	// Compare each job in the response with the expected jobs
	for i := 0; i < len(responseJobs); i++ {
		if responseJobs[i] != jobs[i] {
			t.Errorf("handler returned unexpected job at index %d: got %+v, want %+v", i, responseJobs[i], jobs[i])
		}
	}
}

// TestDeleteJob is a unit test for the deleteJob handler function
func TestDeleteJob(t *testing.T) {
	// Create sample jobs
	jobs := []models.JobRequest{
		{
			JobId:          "1",
			JobName:        "Test",
			JobType:        "jobService",
			JobStatus:      "Open",
			JobPriority:    "High",
			JobDescription: "Job description goes here",
			JobCreatedDate: "2023-04-21",
		},
		{
			JobId:          "2",
			JobName:        "Test2",
			JobType:        "jobService",
			JobStatus:      "Open",
			JobPriority:    "High",
			JobDescription: "Job description goes here",
			JobCreatedDate: "2023-04-22",
		},
	}

	// Convert the job request to JSON
	body, err := json.Marshal(jobs[len(jobs)-1])
	if err != nil {
		t.Fatal(err)
	}

	// Prepare a request with a request body
	req, err := http.NewRequest("DELETE", "/job", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// setup context
	ctx := context.Background()
	// Add the sample job to the job repository mock
	for index := range jobs {
		jobRepo.CreateJob(ctx, &jobs[index])
	}
	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.deleteJob)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the response body matches the expected message
	expectedResponse := "Delete Job!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
	}

	jobsAfterDeleted, err := jobRepo.FindJobs(ctx)
	if err != nil {
		t.Errorf("failed to get jobs after delete: %v", err)
	}
	if len(jobsAfterDeleted) != len(jobs)-1 {
		t.Errorf("failed to delete job: got %d, want %d", len(jobsAfterDeleted), len(jobs)-1)
	}
}

// TestCreateJob is a unit test for the createJob handler function
func TestCreateJob(t *testing.T) {
	// Create a sample job
	job := models.JobRequest{
		JobId:          "1",
		JobName:        "Test",
		JobType:        "jobService",
		JobStatus:      "Open",
		JobPriority:    "High",
		JobDescription: "Job description goes here",
		JobCreatedDate: "2023-04-21",
	}
	// convert job to byte array
	jobJSON, err := json.Marshal(job)
	bodyReader := bytes.NewReader(jobJSON)
	if err != nil {
		t.Errorf("failed to marshal job: %v", err)
	}

	// Prepare a request with a request body
	req, err := http.NewRequest("POST", "/job", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.createJob)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the response body matches the expected message
	expectedResponse := "Create Job!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
	}
	// setup context
	ctx := context.Background()
	// Check if the job was added to the job repository mock
	createdJob, err := jobRepo.FindJobByID(ctx, job.JobId)
	if err != nil {
		t.Errorf("failed to get job: %v", err)
	}
	if *createdJob != job {
		t.Errorf("failed to create job: got %v, want %v", *createdJob, &job)
	}
}

// TestUpdateJob is a unit test for the updateJob handler function
func TestUpdateJob(t *testing.T) {
	// create a sample job
	job := models.JobRequest{
		JobId:          "1",
		JobName:        "Test",
		JobType:        "jobService",
		JobStatus:      "Open",
		JobPriority:    "High",
		JobDescription: "Job description goes here",
		JobCreatedDate: "2023-04-21",
	}
	updateJob := models.JobRequest{
		JobId:          "1",
		JobName:        "Test",
		JobType:        "jobService",
		JobStatus:      "Closed",
		JobPriority:    "High",
		JobDescription: "Job description goes here",
		JobCreatedDate: "2023-04-21",
	}
	// convert job to byte array
	jobJSON, err := json.Marshal(updateJob)
	bodyReader := bytes.NewReader(jobJSON)
	if err != nil {
		t.Errorf("failed to marshal job: %v", err)
	}
	// Prepare a request with a request body
	req, err := http.NewRequest("POST", "/job", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()
	// Create a new job repository mock
	jobRepo := repositories.NewJobsMemory()
	// setup context
	ctx := context.Background()
	// Add the sample job to the job repository mock
	jobRepo.CreateJob(ctx, &job)

	// Create a new server using the job repository mock
	server := NewServer(jobRepo)

	// Create an HTTP handler function from the healthCheck handler
	handler := http.HandlerFunc(server.updateJob)

	// Serve the HTTP request using the handler function
	handler.ServeHTTP(rr, req)

	// Check if the status code of the response is as expected (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if the response body matches the expected message
	expectedResponse := "Update Job!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
	}

	// check if the job was updated to the job repository mock
	updatedJob, err := jobRepo.FindJobByID(ctx, job.JobId)
	if err != nil {
		t.Errorf("failed to get job: %v", err)
	}
	if *updatedJob != updateJob {
		t.Errorf("failed to update job: got %v, want %v", *updatedJob, &updateJob)
	}
}
