package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
)

func (config *Config) session() *gocql.Session {
	cluster := gocql.NewCluster(config.Seeds...)
	cluster.Timeout = time.Second * 3
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 5,
	}

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

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	return session
}
