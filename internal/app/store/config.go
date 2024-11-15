package store

// Config
type Config struct {
	DatabaseURl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
