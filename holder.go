package cqllock

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

// DefaultHolder returns s if it isn't an empty string,
// otherwise it returns the system's hostname.
func DefaultHolder(s string) (holder string) {
	if s != "" {
		return s
	}
	holder, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return
}
