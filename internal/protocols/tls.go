package protocols

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/wybiral/hookah/pkg/node"
)

// TLS creates a TLS client node
func TLS(arg string) (*node.Node, error) {
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
	cfg := &tls.Config{}
	// Handle cert option
	cert := opts.Get("cert")
	if len(cert) != 0 {
		pem, err := ioutil.ReadFile(cert)
		if err != nil {
			return nil, err
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(pem)
		cfg.RootCAs = certPool
	}
	// Handle insecure option
	insecure := opts.Get("insecure")
	if insecure == "true" {
		cfg.InsecureSkipVerify = true
	} else if insecure != "false" && len(insecure) > 0 {
		return nil, errors.New("invalid option for insecure")
	}
	// Open connection with config
	rwc, err := tls.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}
	return node.New(rwc), nil
}
