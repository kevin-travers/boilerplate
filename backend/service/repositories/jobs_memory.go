package repositories

import (
	"context"
	"log"
	"service/models"
)

// fake database using a map of JobRequest structs
type JobsMemory struct {
	JobRequests map[string]*models.JobRequest
}

// NewJobsMemory returns a new JobsMemory struct
func NewJobsMemory() *JobsMemory {
	return &JobsMemory{
		JobRequests: map[string]*models.JobRequest{},
	}
}

// FindJobByID returns a JobRequest struct from the map
func (db *JobsMemory) FindJobByID(ctx context.Context, jobID string) (*models.JobRequest, error) {
	job, ok := db.JobRequests[jobID]
	if !ok {
		log.Printf("Job not found: %s", jobID)
		return nil, models.ErrJobNotFound
	}
	return job, nil
}

// FindJobs returns a slice of JobRequest structs from the map
func (db *JobsMemory) FindJobs(ctx context.Context) ([]*models.JobRequest, error) {
	jobs := []*models.JobRequest{}
	for _, job := range db.JobRequests {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// CreateJob adds a JobRequest struct to the map
func (db *JobsMemory) CreateJob(ctx context.Context, job *models.JobRequest) error {
	db.JobRequests[job.JobId] = job
	return nil
}

// DeleteJob removes a JobRequest struct from the map
func (db *JobsMemory) DeleteJob(ctx context.Context, jobID string) error {
	_, ok := db.JobRequests[jobID]
	if !ok {
		log.Printf("Job not found: %s", jobID)
		return models.ErrJobNotFound
	}
	delete(db.JobRequests, jobID)
	return nil
}

// UpdateJob updates a JobRequest struct in the map
func (db *JobsMemory) UpdateJob(ctx context.Context, job *models.JobRequest) error {
	_, ok := db.JobRequests[job.JobId]
	if !ok {
		log.Printf("Job not found: %s", job.JobId)
		return models.ErrJobNotFound
	}
	db.JobRequests[job.JobId] = job
	return nil
}
