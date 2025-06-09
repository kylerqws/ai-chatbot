package jobstatus

import "github.com/kylerqws/chatbot/pkg/openai/utils/enumset"

// JobStatus defines a fine‑tuning job state.
type JobStatus struct {
	Code        string // Unique identifier for the job status.
	Description string // Human-readable explanation of the job status.
}

// JobStatus code constants.
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
	ValidatingFiles = &JobStatus{Code: ValidatingFilesCode, Description: "Job validation in progress."}
	Queued          = &JobStatus{Code: QueuedCode, Description: "Job is queued."}
	Running         = &JobStatus{Code: RunningCode, Description: "Job is in progress."}
	Succeeded       = &JobStatus{Code: SucceededCode, Description: "Job completed successfully."}
	Cancelled       = &JobStatus{Code: CancelledCode, Description: "Job was cancelled."}
	Failed          = &JobStatus{Code: FailedCode, Description: "Job failed due to error."}
)

// AllJobStatuses lists all known JobStatus instances.
var AllJobStatuses = map[string]*JobStatus{
	ValidatingFilesCode: ValidatingFiles,
	QueuedCode:          Queued,
	RunningCode:         Running,
	SucceededCode:       Succeeded,
	CancelledCode:       Cancelled,
	FailedCode:          Failed,
}

// Resolve looks up a JobStatus by code, error if missing or unknown.
func Resolve(code string) (*JobStatus, error) {
	return enumset.ResolveRequired(code, AllJobStatuses, "job status")
}

// JoinCodes returns all job status codes joined by separator.
func JoinCodes(sep string) string {
	return enumset.JoinCodes(AllJobStatuses, sep)
}
