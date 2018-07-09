package app

// DefaultBufferSize is default internal buffer size for readers.
const DefaultBufferSize = 8 * 1024

// Config options for App.
type Config struct {
	BufferSize int
}
