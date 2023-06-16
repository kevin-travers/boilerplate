package search

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
)

type Search struct {
	client   *elasticsearch.Client
	jobIndex string
}

func NewSearch(elasticSearchURL string, jobIndex string) (*Search, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticSearchURL,
		},
		Username: "elastic",
		Password: "password",
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Search{
		client:   client,
		jobIndex: jobIndex,
	}, nil
}

// setup index if does not exists
func (s *Search) SetupIndex(ctx context.Context, name string) error {
	res, err := s.client.Indices.Exists([]string{name})
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		_, err := s.client.Indices.Create(name)
		if err != nil {
			return err
		}
	}
	return nil
}

// create job into index
func (s *Search) CreateJob(ctx context.Context, job *jobSearchMapping) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(job); err != nil {
		return err
	}
	res, err := s.client.Index(
		s.jobIndex,
		&buf,
		s.client.Index.WithDocumentID(job.JobId),
		s.client.Index.WithRefresh("true"),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

//serach jobs
//update jobs
//bulk search jobs
//job to mapping
