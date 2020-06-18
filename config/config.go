package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Server string
	Name   string
}

func NewConfig(server, name string) Config {
	return Config{
		DB: DBConfig{
			Server: server,
			Name:   name,
		},
	}
}
