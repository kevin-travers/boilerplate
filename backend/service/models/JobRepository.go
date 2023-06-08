package models

type JobRepository interface {
	FindJobByID(jobID string) (*JobRequest, error)
	FindJobs() ([]*JobRequest, error)
	CreateJob(job *JobRequest) error
	UpdateJob(job *JobRequest) error
	DeleteJob(jobID string) error
}
