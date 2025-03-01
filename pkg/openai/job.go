package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	jobCreateURL = "https://api.openai.com/v1/jobs"
	jobInfoURL   = "https://api.openai.com/v1/jobs/%s"
	jobCancelURL = "https://api.openai.com/v1/jobs/%s/cancel"
)

// JobClient управляет заданиями в OpenAI API.
type JobClient struct {
	apiKey string
	client *http.Client
}

// NewJobClient создаёт новый экземпляр Job API Client.
func NewJobClient(apiKey string) *JobClient {
	return &JobClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// CreateJob создаёт новое задание.
func (jc *JobClient) CreateJob(params map[string]interface{}) (string, error) {
	requestBody, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", jobCreateURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create job request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+jc.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := jc.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("job creation failed: %v", err)
	}
	defer resp.Body.Close()

	var response struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode job creation response: %v", err)
	}

	return response.ID, nil
}

// GetJobInfo получает информацию о задании по его ID.
func (jc *JobClient) GetJobInfo(jobID string) (map[string]interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf(jobInfoURL, jobID), nil)
	req.Header.Set("Authorization", "Bearer "+jc.apiKey)

	resp, err := jc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get job info: %v", err)
	}
	defer resp.Body.Close()

	var jobInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jobInfo); err != nil {
		return nil, fmt.Errorf("failed to decode job info response: %v", err)
	}

	return jobInfo, nil
}

// CancelJob отменяет задание по его ID.
func (jc *JobClient) CancelJob(jobID string) error {
	req, _ := http.NewRequest("POST", fmt.Sprintf(jobCancelURL, jobID), nil)
	req.Header.Set("Authorization", "Bearer "+jc.apiKey)

	resp, err := jc.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to cancel job: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

// ListJobs получает список всех заданий.
func (jc *JobClient) ListJobs() ([]map[string]interface{}, error) {
	req, _ := http.NewRequest("GET", jobCreateURL, nil)
	req.Header.Set("Authorization", "Bearer "+jc.apiKey)

	resp, err := jc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list jobs: %v", err)
	}
	defer resp.Body.Close()

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode list jobs response: %v", err)
	}

	return response.Data, nil
}

// CancelAllJobs отменяет все задания.
func (jc *JobClient) CancelAllJobs() error {
	jobs, err := jc.ListJobs()
	if err != nil {
		return err
	}

	for _, job := range jobs {
		jobID, ok := job["id"].(string)
		if !ok {
			continue
		}
		_ = jc.CancelJob(jobID)
	}

	return nil
}
