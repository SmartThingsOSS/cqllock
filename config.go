package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// Config contains the Cassandra configuration required for the
// cqllock tools to run.
type Config struct {
	Seeds    []string
	CertPath string
	KeyPath  string
	Username string
	Password string
	Keyspace string
	Table    string
	Timeout  int
	Retries  int
}

const configFile = "~/.cqllockrc"

// ParseConfig parses the cqllock config file into a Config object.
func parseConfig() *Config {
	config := Config{}
	contents, err := ioutil.ReadFile(expandHome(configFile))
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal([]byte(contents), &config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func expandHome(path string) (ret string) {
	ret, err := homedir.Expand(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}
