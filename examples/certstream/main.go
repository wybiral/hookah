// Example of using hookah to create an input stream from the CertStream
// WebSocket API (https://certstream.calidog.io/).
// The cert updates are filtered to remove heartbeat messages and processed by
// restricting the JSON fields and adding indentation.
// These updates are then written to stdout.
package main

import (
	"encoding/json"
	"log"

	"github.com/wybiral/hookah"
	"github.com/wybiral/hookah/pkg/node"
)

// CertStream JSON struct
type certUpdate struct {
	MessageType string `json:"message_type"`
	Data        struct {
		UpdateType string `json:"update_type"`
		LeafCert   struct {
			AllDomains  []string `json:"all_domains"`
			Fingerprint string   `json:"fingerprint"`
		} `json:"leaf_cert"`
		Source struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		}
	} `json:"data"`
}

func main() {
	// Create hookah API instance
	h := hookah.New()
	// Create hookah input (certstream WebSocket API)
	r, err := h.NewNode("wss://certstream.calidog.io")
	if err != nil {
		log.Fatal(err)
	}
	// Create hookah output (stdout)
	w, err := h.NewNode("stdout")
	if err != nil {
		log.Fatal(err)
	}
	// Start stream
	stream(w, r)
}

// Copy from reader to writer
// Drops heartbeat messages, restricts fields, and formats JSON
func stream(w, r *node.Node) {
	var u certUpdate
	d := json.NewDecoder(r.Reader())
	e := json.NewEncoder(w.Writer())
	e.SetIndent("", "  ")
	for {
		err := d.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
		if u.MessageType == "heartbeat" {
			continue
		}
		err = e.Encode(u.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
