package jobstatus

import (
	"fmt"
	"strings"
)

type JobStatus struct {
	Code        string
	Description string
}

const (
	ValidatingCode = "validating"
	RunningCode    = "running"
	SucceededCode  = "succeeded"
	CancelledCode  = "cancelled"
	FailedCode     = "failed"
)

var (
	Validating = &JobStatus{
		Code:        ValidatingCode,
		Description: "The job is currently being validated.",
	}
	Running = &JobStatus{
		Code:        RunningCode,
		Description: "The job is currently running.",
	}
	Succeeded = &JobStatus{
		Code:        SucceededCode,
		Description: "The job completed successfully.",
	}
	Cancelled = &JobStatus{
		Code:        CancelledCode,
		Description: "The job was cancelled.",
	}
	Failed = &JobStatus{
		Code:        FailedCode,
		Description: "The job failed to complete.",
	}
)

var AllJobStatuses = map[string]*JobStatus{
	ValidatingCode: Validating,
	RunningCode:    Running,
	SucceededCode:  Succeeded,
	CancelledCode:  Cancelled,
	FailedCode:     Failed,
}

func Resolve(code string) (*JobStatus, error) {
	if status, ok := AllJobStatuses[code]; ok {
		return status, nil
	}

	return nil, fmt.Errorf("unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllJobStatuses))
	for code := range AllJobStatuses {
		codes = append(codes, code)
	}

	return strings.Join(codes, sep)
}
