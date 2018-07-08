package app

const DefaultBufferSize = 8 * 1024

type Config struct {
	RWOpts     []string
	ROpts      []string
	WOpts      []string
	BufferSize int
}
