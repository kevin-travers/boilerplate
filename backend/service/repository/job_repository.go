package repository

import (
	"service/models"
)

type JobRepository interface {
	GetJobByID(id string) (*models.JobRequest, error)
	GetJobs() ([]*models.JobRequest, error)
	CreateJob(job *models.JobRequest) (*models.JobRequest, error)
	DeleteJob(id string) error
	UpdateJob(job *models.JobRequest) error
}
