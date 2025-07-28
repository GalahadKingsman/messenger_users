package config

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type Config struct {
	DB       DBConfig
	GRPCPort int `env:"GRPC_PORT"`
}
