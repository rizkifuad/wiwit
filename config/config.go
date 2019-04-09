package config

type Config struct {
	DbHost string
}

func New() Config {
	return Config{DbHost: "menggila"}
}
