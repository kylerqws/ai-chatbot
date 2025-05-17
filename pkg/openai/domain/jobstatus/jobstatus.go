package jobstatus

import (
	"fmt"
	"strings"
)

type JobStatus struct {
	Code        string
	Description string
}

var (
	Validating = JobStatus{
		Code:        "validating",
		Description: "The job is currently being validated.",
	}
	Running = JobStatus{
		Code:        "running",
		Description: "The job is currently running.",
	}
	Succeeded = JobStatus{
		Code:        "succeeded",
		Description: "The job completed successfully.",
	}
	Failed = JobStatus{
		Code:        "failed",
		Description: "The job failed to complete.",
	}
	Cancelled = JobStatus{
		Code:        "cancelled",
		Description: "The job was cancelled.",
	}
)

var AllJobStatuses = map[string]*JobStatus{
	Validating.Code: &Validating,
	Running.Code:    &Running,
	Succeeded.Code:  &Succeeded,
	Failed.Code:     &Failed,
	Cancelled.Code:  &Cancelled,
}

func Resolve(code string) (*JobStatus, error) {
	if sts, ok := AllJobStatuses[code]; ok {
		return sts, nil
	}

	return nil, fmt.Errorf("unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	var codes []string
	for _, sts := range AllJobStatuses {
		codes = append(codes, sts.Code)
	}

	return strings.Join(codes, sep)
}
