package repositories

import "service/models"

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

// NewJobsMemoryWithJobs returns a new JobsMemory struct with some prepopulated jobs
func NewJobsMemoryWithJobs() *JobsMemory {
	db := NewJobsMemory()
	db.JobRequests["1"] = &models.JobRequest{
		JobId:          "1",
		JobName:        "Job 1",
		JobType:        "Job Type 1",
		JobStatus:      "Job Status 1",
		JobPriority:    "Job Priority 1",
		JobDescription: "Job Description 1",
		JobCreatedDate: "Job Created Date 1",
		JobUpdatedDate: "Job Updated Date 1",
	}
	db.JobRequests["2"] = &models.JobRequest{
		JobId:          "2",
		JobName:        "Job 2",
		JobType:        "Job Type 2",
		JobStatus:      "Job Status 2",
		JobPriority:    "Job Priority 2",
		JobDescription: "Job Description 2",
		JobCreatedDate: "Job Created Date 2",
		JobUpdatedDate: "Job Updated Date 2",
	}
	db.JobRequests["3"] = &models.JobRequest{
		JobId:          "3",
		JobName:        "Job 3",
		JobType:        "Job Type 3",
		JobStatus:      "Job Status 3",
		JobPriority:    "Job Priority 3",
		JobDescription: "Job Description 3",
		JobCreatedDate: "Job Created Date 3",
		JobUpdatedDate: "Job Updated Date 3",
	}
	return db
}

// FindJobByID returns a JobRequest struct from the map
func (db *JobsMemory) FindJobByID(jobID string) (*models.JobRequest, error) {
	job, ok := db.JobRequests[jobID]
	if !ok {
		return nil, models.ErrJobNotFound
	}
	return job, nil
}

// GetJobs returns a slice of JobRequest structs from the map
func (db *JobsMemory) GetJobs() ([]*models.JobRequest, error) {
	jobs := []*models.JobRequest{}
	for _, job := range db.JobRequests {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// CreateJob adds a JobRequest struct to the map
func (db *JobsMemory) CreateJob(job *models.JobRequest) error {
	db.JobRequests[job.JobId] = job
	return nil
}

// DeleteJob removes a JobRequest struct from the map
func (db *JobsMemory) DeleteJob(jobID string) error {
	_, ok := db.JobRequests[jobID]
	if !ok {
		return models.ErrJobNotFound
	}
	delete(db.JobRequests, jobID)
	return nil
}

// UpdateJob updates a JobRequest struct in the map
func (db *JobsMemory) UpdateJob(job *models.JobRequest) error {
	_, ok := db.JobRequests[job.JobId]
	if !ok {
		return models.ErrJobNotFound
	}
	db.JobRequests[job.JobId] = job
	return nil
}
