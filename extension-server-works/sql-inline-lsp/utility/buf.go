package utility

import (
	"io"
	"os"
)

type bufWrapper struct {
	reader io.Reader
	writer io.Writer
}

// Close implements io.ReadWriteCloser.
func (b *bufWrapper) Close() error {
	return nil
}

// Read implements io.ReadWriteCloser.
func (b *bufWrapper) Read(p []byte) (n int, err error) {
	return b.reader.Read(p)
}

// Write implements io.ReadWriteCloser.
func (b *bufWrapper) Write(p []byte) (n int, err error) {
	return b.writer.Write(p)
}

var _ io.ReadWriteCloser = (*bufWrapper)(nil)

func CreateStdReadWriteCloser() io.ReadWriteCloser {
	bf := &bufWrapper{
		reader: os.Stdin,
		writer: os.Stdout,
	}

	return bf
}
