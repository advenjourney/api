package config

// Config describe a general configuration struct for the application
type Config struct {
	Database DB
	Server   Server
}

// DB describes the database config object
type DB struct {
	DSN string
}

type Server struct {
	Addr string
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{
		Server: Server{
			Addr: ":8080",
		},
	}
}
