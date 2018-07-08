package app

// DefaultBufferSize is default internal buffer size for readers.
const DefaultBufferSize = 8 * 1024

// Config options for App.
type Config struct {
	RWOpts     []string
	ROpts      []string
	WOpts      []string
	BufferSize int
}
