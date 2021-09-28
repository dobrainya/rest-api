package store

// Config ...
type Config struct {
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
