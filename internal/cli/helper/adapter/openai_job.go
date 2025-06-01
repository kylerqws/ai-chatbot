package adapter

import (
	"fmt"
	"regexp"

	"github.com/kylerqws/chatbot/pkg/openai/contract/service"
	"github.com/spf13/cobra"
)

type Job struct {
	*service.Job
	ExecStatus bool
}

type OpenAiJobAdapter struct {
	command *cobra.Command
	jobs    []*Job
}

func NewOpenAiJobAdapter(cmd *cobra.Command) *OpenAiJobAdapter {
	return &OpenAiJobAdapter{command: cmd}
}

func (h *OpenAiJobAdapter) Jobs() []*Job {
	return h.jobs
}

func (h *OpenAiJobAdapter) ExistJobs() bool {
	return len(h.jobs) > 0
}

func (*OpenAiJobAdapter) WrapOpenAIJob(job *service.Job) *Job {
	return &Job{Job: job}
}

func (h *OpenAiJobAdapter) WrapOpenAIJobs(jobs ...*service.Job) []*Job {
	wraps := make([]*Job, len(jobs))
	for i := range jobs {
		wraps = append(wraps, h.WrapOpenAIJob(jobs[i]))
	}
	return wraps
}

func (h *OpenAiJobAdapter) AddJob(job *Job) {
	if job != nil {
		h.jobs = append(h.jobs, job)
	}
}

func (h *OpenAiJobAdapter) AddJobs(jobs ...*Job) {
	for i := range jobs {
		h.AddJob(jobs[i])
	}
}

func (*OpenAiJobAdapter) ValidateJobID(jobID string) error {
	if !regexp.MustCompile(`^ftjob-[a-zA-Z0-9]{24,}$`).MatchString(jobID) {
		return fmt.Errorf("invalid job ID '%s'", jobID)
	}
	return nil
}
