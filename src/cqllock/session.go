package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
)

func (config *Config) session() *gocql.Session {
	cluster := gocql.NewCluster(config.Seeds...)

	if config.CertPath != "" {
		log.Debugf("using SSL with certificate '%s' and key '%s'", config.CertPath, config.KeyPath)
		cluster.SslOpts = &gocql.SslOptions{
			CertPath:               expandHome(config.CertPath),
			KeyPath:                expandHome(config.KeyPath),
			EnableHostVerification: false,
		}
	}

	if config.Username != "" {
		log.Debugf("using authentication with username '%s'", config.Username)
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Username,
			Password: config.Password,
		}
	}

	if config.Timeout > 0 {
		cluster.Timeout = time.Second * time.Duration(config.Timeout)
		log.Debugf("setting cluster timeout to %v", cluster.Timeout)
	}

	if config.Retries > 0 {
		log.Debugf("setting cluster retry number to %d", config.Retries)
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
