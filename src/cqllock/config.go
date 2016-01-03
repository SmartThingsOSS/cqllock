package main

import (
	"errors"
	"io/ioutil"
	"os"

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

var configFiles = []string{"~/.cqllockrc", "/etc/cqllock.yaml"}

// ParseConfig parses the cqllock config file into a Config object.
func parseConfig() *Config {
	config := Config{}

	path, err := configPath()
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(contents, &config); err != nil {
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

func configPath() (path string, err error) {
	for _, path = range configFiles {
		path = expandHome(path)
		if _, err = os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path, nil
	}
	return "", errors.New("no config file found")
}
