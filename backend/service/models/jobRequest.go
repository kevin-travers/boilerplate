package models

import (
	"encoding/json"
	"io"
)

type JobRequest struct {
	JobId          string `json:"JobId"`
	JobName        string `json:"JobName"`
	JobType        string `json:"JobType"`
	JobStatus      string `json:"JobStatus"`
	JobPriority    string `json:"JobPriority"`
	JobDescription string `json:"JobDescription"`
	JobCreatedDate string `json:"JobCreatedDate"`
}

type JobRequestFilter struct {
	JobId          string `json:"JobId"`
	JobName        string `json:"JobName"`
	JobType        string `json:"JobType"`
	JobStatus      string `json:"JobStatus"`
	JobPriority    string `json:"JobPriority"`
	JobDescription string `json:"JobDescription"`
	JobCreatedDate string `json:"JobCreatedDate"`
	Offset         int    `json:"Offset"`
	Limit          int    `json:"Limit"`
}

func (job *JobRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(job)
}

func (job *JobRequestFilter) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(job)
}
