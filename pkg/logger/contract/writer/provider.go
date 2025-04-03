package writer

import "io"

type Provider interface {
	Writer() io.Writer
}
