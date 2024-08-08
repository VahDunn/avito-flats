package config

type Config struct {
	ServerAddress      string
	DBConnectionString string
}

func LoadConfig() *Config {
	// Загрузка конфигурации
	return &Config{
		ServerAddress: ":8080",
		// todo убрать секреты в переменные окружения
		DBConnectionString: "user =avitodev password=pg4afl dbname=avito-flats sslmode=disable host=localhostport=5432",
	}
}
