package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
)

func (config *Config) session() *gocql.Session {
	cluster := gocql.NewCluster(config.Seeds...)

	if config.CertPath != "" {
		cluster.SslOpts = &gocql.SslOptions{
			CertPath:               expandHome(config.CertPath),
			KeyPath:                expandHome(config.KeyPath),
			EnableHostVerification: false,
		}
	}

	if config.Username != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Username,
			Password: config.Password,
		}
	}

	if config.Timeout > 0 {
		cluster.Timeout = time.Second * config.Timeout
	}

	if config.Retries > 0 {
		cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
			NumRetries: config.Retries,
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	return session
}
