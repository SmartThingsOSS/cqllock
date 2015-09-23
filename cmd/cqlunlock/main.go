package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	log "github.com/Sirupsen/logrus"
	"github.com/SmartThingsOSS/cqllock"
	"github.com/SmartThingsOSS/stcql"
)

var (
	lockName   = kingpin.Arg("name", "Name of the mutex to unlock.").Required().String()
	holderFlag = kingpin.Flag("holder", "Name of the lock holder. Defaults to hostname.").Short('h').String()
)

func main() {
	kingpin.Parse()
	holder := cqllock.DefaultHolder(*holderFlag)
	config := cqllock.ParseConfig()

	m := &stcql.Mutex{
		Session:    config.Session(),
		Keyspace:   config.Keyspace,
		Table:      config.Table,
		LockHolder: holder,
		LockName:   *lockName,
	}
	defer m.Session.Close()

	if err := m.Unlock(); err != nil {
		log.Fatalf("failed to unlock '%s'", m.LockName)
	}

	os.Exit(0)
}
