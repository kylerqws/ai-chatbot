package jobstatus

import (
	"fmt"
	"sort"
	"strings"
)

// JobStatus defines the current state of a fine-tuning job.
type JobStatus struct {
	Code        string // Unique identifier for the job status.
	Description string // Human-readable explanation of the job status.
}

// Job status code constants.
const (
	ValidatingFilesCode = "validating_files"
	QueuedCode          = "queued"
	RunningCode         = "running"
	SucceededCode       = "succeeded"
	CancelledCode       = "cancelled"
	FailedCode          = "failed"
)

// Predefined JobStatus instances.
var (
	ValidatingFiles = &JobStatus{
		Code:        ValidatingFilesCode,
		Description: "The job is being validated before execution.",
	}
	Queued = &JobStatus{
		Code:        QueuedCode,
		Description: "The job is waiting in the queue to be processed.",
	}
	Running = &JobStatus{
		Code:        RunningCode,
		Description: "The job is currently in progress.",
	}
	Succeeded = &JobStatus{
		Code:        SucceededCode,
		Description: "The job finished successfully.",
	}
	Cancelled = &JobStatus{
		Code:        CancelledCode,
		Description: "The job was cancelled by the user or system.",
	}
	Failed = &JobStatus{
		Code:        FailedCode,
		Description: "The job failed to complete due to an error.",
	}
)

// AllJobStatuses maps all known job status codes to their JobStatus instances.
var AllJobStatuses = map[string]*JobStatus{
	ValidatingFilesCode: ValidatingFiles,
	QueuedCode:          Queued,
	RunningCode:         Running,
	SucceededCode:       Succeeded,
	CancelledCode:       Cancelled,
	FailedCode:          Failed,
}

// Resolve returns the JobStatus associated with the given code.
// Returns an error if the code is empty or unrecognized.
func Resolve(code string) (*JobStatus, error) {
	if code == "" {
		return nil, fmt.Errorf("job status code is required")
	}
	if jst, ok := AllJobStatuses[code]; ok {
		return jst, nil
	}
	return nil, fmt.Errorf("unknown job status code: '%s'", code)
}

// JoinCodes returns a sorted, delimited string of all known job status codes.
func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllJobStatuses))
	for code := range AllJobStatuses {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
