package config

// TODO: Handle database path in configuration
type Config struct {
	DatabaseName string
}

func NewConfig(databaseName string) Config {
	return Config{
		DatabaseName: databaseName,
	}
}
