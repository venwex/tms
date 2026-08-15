package config

type Config struct {
	Postgres PostgresConfig
}

func New() *Config {
	return &Config{
		Postgres: NewPostgresConfig(),
	}
}
