package filestatus

import (
	"fmt"
	"strings"
)

type FileStatus struct {
	Code        string
	Description string
}

const (
	UploadedCode  = "uploaded"
	ProcessedCode = "processed"
	DeletedCode   = "deleted"
	ErrorCode     = "error"
)

var (
	Uploaded = &FileStatus{
		Code:        UploadedCode,
		Description: "The file has been successfully uploaded and is awaiting processing.",
	}
	Processed = &FileStatus{
		Code:        ProcessedCode,
		Description: "The file has been successfully processed and is ready for use.",
	}
	Deleted = &FileStatus{
		Code:        DeletedCode,
		Description: "The file has been deleted and is no longer available.",
	}
	Error = &FileStatus{
		Code:        ErrorCode,
		Description: "An error occurred during file processing.",
	}
)

var AllFileStatuses = map[string]*FileStatus{
	UploadedCode:  Uploaded,
	ProcessedCode: Processed,
	DeletedCode:   Deleted,
	ErrorCode:     Error,
}

func Resolve(code string) (*FileStatus, error) {
	if status, ok := AllFileStatuses[code]; ok {
		return status, nil
	}

	return nil, fmt.Errorf("unknown value '%v'", code)
}

func JoinCodes(sep string) string {
	codes := make([]string, 0, len(AllFileStatuses))
	for code := range AllFileStatuses {
		codes = append(codes, code)
	}

	return strings.Join(codes, sep)
}
