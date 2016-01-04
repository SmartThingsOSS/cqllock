package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	log "github.com/Sirupsen/logrus"
	"github.com/SmartThingsOSS/stcql"
)

var (
	lockName   = kingpin.Arg("name", "Name of the lock to acquire.").Required().String()
	retryTime  = kingpin.Flag("retry-time", "Time between retries when acquiring lock. 0 means don't retry.").Short('r').Default("0").Duration()
	timeout    = kingpin.Flag("timeout", "Timeout when retrying acquiring lock. 0 means try forever.").Short('t').Default("0").Duration()
	lifetime   = kingpin.Flag("lifetime", "Lifetime of the lock. Lock will be considered stale after this duration. 0 means lock is valid forever.").Short('l').Default("0").Duration()
	holderFlag = kingpin.Flag("holder", "Name of the lock holder. Defaults to hostname.").Short('h').String()
	unlock     = kingpin.Flag("unlock", "Unlock instead of lock.").Short('u').Default("false").Bool()
	debug      = kingpin.Flag("debug", "Print debugging messages.").Default("false").Bool()
	version    = "dev"
)

func defaultHolder() *string {
	if *holderFlag != "" {
		return holderFlag
	}
	holder, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return &holder
}

func main() {
	kingpin.Version(version)
	kingpin.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	// set holder to hostname if it's not specified on the command-line
	holder := defaultHolder()
	config := parseConfig()

	m := &stcql.Mutex{
		Session:    config.session(),
		Keyspace:   config.Keyspace,
		Table:      config.Table,
		LockHolder: *holder,
		Lifetime:   *lifetime,
		LockName:   *lockName,
		RetryTime:  *retryTime,
		Timeout:    *timeout,
	}
	defer m.Session.Close()

	if !*unlock {
		if err := m.Lock(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := m.Unlock(); err != nil {
			log.Fatalf("failed to unlock '%s'", m.LockName)
		}
	}

	os.Exit(0)
}
