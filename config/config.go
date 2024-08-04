package config

type Config struct {
	ServerAddress      string
	DBConnectionString string
}

func LoadConfig() *Config {
	// Загрузка конфигурации
	return &Config{
		ServerAddress:      ":8080",
		DBConnectionString: "user=your_user password=your_password dbname=your_db sslmode=disable",
	}
}
