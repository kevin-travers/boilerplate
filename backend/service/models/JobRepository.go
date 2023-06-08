package models

import "context"

type JobRepository interface {
	FindJobByID(ctx context.Context, jobID string) (*JobRequest, error)
	FindJobs(ctx context.Context) ([]*JobRequest, error)
	CreateJob(ctx context.Context, job *JobRequest) error
	UpdateJob(ctx context.Context, job *JobRequest) error
	DeleteJob(ctx context.Context, jobID string) error
}
