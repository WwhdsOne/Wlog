package file

import "io"

type Writer interface {
	CreateWriter() io.Writer
}
