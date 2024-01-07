package config

func New() *Config {
	cfg := &Config{}

	// TODO: don't hardcode this.
	cfg.API.Port = ":9000"

	return cfg
}
