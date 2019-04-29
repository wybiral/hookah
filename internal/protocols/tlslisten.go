package protocols

import (
	"crypto/tls"
	"net/url"
	"strings"

	"github.com/wybiral/hookah/pkg/fanout"
	"github.com/wybiral/hookah/pkg/node"
)

// TLSListen creates a TLS listener Node
func TLSListen(arg string) (*node.Node, error) {
	var opts url.Values
	// Parse options
	addrOpts := strings.SplitN(arg, "?", 2)
	addr := addrOpts[0]
	if len(addrOpts) == 2 {
		op, err := url.ParseQuery(addrOpts[1])
		if err != nil {
			return nil, err
		}
		opts = op
	}
	// Load cert and key files
	cert := opts.Get("cert")
	key := opts.Get("key")
	app := &listenApp{}
	c, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{c}}
	// Create listener from config
	ln, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}
	app.ln = ln
	app.fan = fanout.New()
	app.ch = make(chan []byte)
	go app.serve()
	return node.New(app), nil
}
