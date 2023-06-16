package search

// JobSearchMapping is json mapping for the job that can be searched in elastic search
type jobSearchMapping struct {
	JobId          string `json:"JobId"`
	JobName        string `json:"JobName"`
	JobType        string `json:"JobType"`
	JobStatus      string `json:"JobStatus"`
	JobPriority    string `json:"JobPriority"`
	JobCreatedDate string `json:"JobCreatedDate"`
}
