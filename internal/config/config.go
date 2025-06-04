package config

type DBConfig struct {
	// Database settings
	DBHost     string `env:"HOST"`
	DBPort     string `env:"PORT"`
	DBUser     string `env:"USER"`
	DBPassword string `env:"PASSWORD"`
	DBName     string `env:"NAME"`
}

type Config struct {
	DBConfig DBConfig `env_prefix:"DB_"`
	GRPCPort int      `env:"GRPC_PORT" envDefault:"9000"`
}

//var (
//	once     sync.Once
//	instance *Config
//)

// getConfig возвращает экземпляр конфигурации singleton
//func GetConfig() *Config {
//	once.Do(func() {
//		instance = &Config{
//			// Database settings
//			DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
//			DBPort:     getEnvOrDefault("DB_PORT", "5432"),
//			DBUser:     getEnvOrDefault("DB_USER", "postgres"),
//			DBPassword: getEnvOrDefault("DB_PASSWORD", "qwerty"),
//			DBName:     getEnvOrDefault("DB_NAME", "messenger_users"),
//
//			// Добавьте другие настройки по мере необходимости
//		}
//	})
//	return instance
//}

// Вспомогательная функция для получения переменной окружения со значением по умолчанию
//func getEnvOrDefault(key, defaultValue string) string {
//	if value, exists := os.LookupEnv(key); exists {
//		return value
//	}
//	return defaultValue
//}
