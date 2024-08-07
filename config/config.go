package config

type Config struct {
	ServerAddress      string
	DBConnectionString string
}

func LoadConfig() *Config {
	// Загрузка конфигурации
	return &Config{
		ServerAddress:      ":8080",
		DBConnectionString: "user=avitodev password=pg4afl dbname=avito-flats sslmode=disable host=localhost port=5432",
	}
}
