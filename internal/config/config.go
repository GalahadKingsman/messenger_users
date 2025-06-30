package config

// Config содержит всю конфигурацию приложения
type DBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME" envDefault:"messenger_users"`
}

type Config struct {
	DB       DBConfig `envPrefix:"DB_"`
	GRPCPort int      `env:"GRPC_PORT" envDefault:"9000"`
}

//var (
//	once     sync.Once
//	instance *Config
//)
//
//// getConfig возвращает экземпляр конфигурации singleton
//func GetConfig(config DBConfig) *Config {
//	once.Do(func() {
//		instance = &Config{
//			// Database settings
//			Host:     getEnvOrDefault("DB_HOST", "localhost"),
//			Port:     getEnvOrDefault("DB_PORT", "5432"),
//			User:     getEnvOrDefault("DB_USER", "postgres"),
//			Password: getEnvOrDefault("DB_PASSWORD", "qwerty"),
//			Name:     getEnvOrDefault("DB_NAME", "messenger_users"),
//
//			// Добавьте другие настройки по мере необходимости
//		}
//	})
//	return instance
//}
//
//// Вспомогательная функция для получения переменной окружения со значением по умолчанию
//func getEnvOrDefault(key, defaultValue string) string {
//	if value, exists := os.LookupEnv(key); exists {
//		return value
//	}
//	return defaultValue
//}
