package writer

import "io"

// Writer types used to identify logging destinations.
const (
	TypeStdout = "stdout"
	TypeStderr = "stderr"
	TypeDB     = "db"
)

// Provider provides an io.Writer for logger output.
type Provider interface {
	// Writer returns the configured io.Writer.
	Writer() io.Writer
}
