package config

import "os"

func New() *Config {
	cfg := &Config{}

	// TODO: don't hardcode this.
	cfg.API.Port = ":9000"

	// TODO: determine where and how to store config.
	// TODO: use different connection for master and slave.
	cfg.Database.Master = os.Getenv("DATABASE_URL")
	cfg.Database.Slave = os.Getenv("DATABASE_URL")

	return cfg
}
