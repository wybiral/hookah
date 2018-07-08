package protocols

import (
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/wybiral/hookah/pkg/node"
)

// File creates a file input and returns Node
func File(path string, opts url.Values) (node.Node, error) {
	perm := os.FileMode(0666)
	permstr := opts.Get("perm")
	if len(permstr) > 0 {
		p, err := strconv.ParseInt(permstr, 10, 32)
		if err != nil {
			return nil, err
		}
		perm = os.FileMode(p)
	}
	flags := os.O_CREATE
	mode := "rwa"
	m := opts.Get("mode")
	if len(m) > 0 {
		mode = m
	}
	if strings.Contains(mode, "a") {
		flags |= os.O_APPEND
	}
	if strings.Contains(mode, "t") {
		flags |= os.O_TRUNC
	}
	read := strings.Contains(mode, "r")
	write := strings.Contains(mode, "w")
	if read && write {
		flags |= os.O_RDWR
	} else if read {
		flags |= os.O_RDONLY
	} else if write {
		flags |= os.O_WRONLY
	}
	return os.OpenFile(path, flags, perm)
}
