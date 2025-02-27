package training

import (
	"fmt"
	"time"
)

// Trainer управляет процессом обучения модели.
type Trainer struct {
	client *OpenAIClient
}

// JobStatus - возможные статусы job.
const (
	StatusRunning  = "running"
	StatusQueued   = "queued"
	StatusFailed   = "failed"
	StatusCanceled = "cancelled"
)

// Job содержит информацию о задаче обучения.
type Job struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// NewTrainer создаёт новый экземпляр Trainer.
func NewTrainer(apiKey string) *Trainer {
	client := NewOpenAIClient(apiKey)
	return &Trainer{client: client}
}

// CreateTrainingJob создаёт задание на обучение модели.
func (t *Trainer) CreateTrainingJob(trainingFileID, model string) (string, error) {
	body := map[string]string{
		"training_file": trainingFileID,
		"model":         model,
		"suffix":        "custom_finetuned_model",
	}

	response, err := t.client.SendRequest("POST", trainURL, body)
	if err != nil {
		return "", fmt.Errorf("failed to create training job: %v", err)
	}

	jobID, exists := response["id"].(string)
	if !exists {
		return "", fmt.Errorf("training job ID not found in response")
	}

	return jobID, nil
}

// GetTrainingJobInfo получает информацию о статусе обучения.
func (t *Trainer) GetTrainingJobInfo(jobID string) (map[string]interface{}, error) {
	url := fmt.Sprintf(jobURL, jobID)
	return t.client.SendRequest("GET", url, nil)
}

// CancelTrainingJob отменяет задание на обучение.
func (t *Trainer) CancelTrainingJob(jobID string) error {
	url := fmt.Sprintf(cancelURL, jobID)
	_, err := t.client.SendRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to cancel training job: %v", err)
	}
	return nil
}

// CancelAllJobs отменяет все активные задания на обучение.
func (t *Trainer) CancelAllJobs() error {
	// Получаем список всех jobs
	response, err := t.client.SendRequest("GET", trainURL, nil)
	if err != nil {
		return fmt.Errorf("failed to fetch jobs list: %v", err)
	}

	jobs, exists := response["data"].([]interface{})
	if !exists {
		return fmt.Errorf("unexpected response format")
	}

	for _, j := range jobs {
		jobData, ok := j.(map[string]interface{})
		if !ok {
			continue
		}

		jobID, _ := jobData["id"].(string)
		status, _ := jobData["status"].(string)

		// Отменяем только активные задания
		if status == "running" || status == "queued" {
			err := t.CancelTrainingJob(jobID)
			if err != nil {
				fmt.Printf("Failed to cancel job %s: %v\n", jobID, err)
			} else {
				fmt.Printf("Job %s canceled successfully\n", jobID)
			}
		}
	}
	return nil
}

// ListJobs получает список jobs, отфильтрованный по статусу и дате.
func (t *Trainer) ListJobs(status string, since time.Time) ([]Job, error) {
	response, err := t.client.SendRequest("GET", trainURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs list: %v", err)
	}

	jobsRaw, exists := response["data"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("unexpected response format")
	}

	var jobs []Job
	for _, j := range jobsRaw {
		jobData, ok := j.(map[string]interface{})
		if !ok {
			continue
		}

		job := Job{
			ID:        jobData["id"].(string),
			Status:    jobData["status"].(string),
			CreatedAt: int64(jobData["created_at"].(float64)),
		}

		// Фильтрация по статусу
		if status != "" && job.Status != status {
			continue
		}

		// Фильтрация по дате
		jobTime := time.Unix(job.CreatedAt, 0)
		if jobTime.Before(since) {
			continue
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}
