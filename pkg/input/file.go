package input

import (
	"io"
	"os"
)

func file(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
