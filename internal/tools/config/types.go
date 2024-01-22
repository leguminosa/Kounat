package config

type Config struct {
	API      ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Master string
	Slave  string
}
