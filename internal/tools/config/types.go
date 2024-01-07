package config

type Config struct {
	API ServerConfig
}

type ServerConfig struct {
	Port string
}
