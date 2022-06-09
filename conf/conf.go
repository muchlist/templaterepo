package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User    UserConfig
	Complex ComplexConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		// dont panic, because in prod we will not use this
		// use os env instead
		log.Print(".env notfound")
	}

	return &Config{
		User: UserConfig{
			ServerPort: os.Getenv("SERVER_PORT"),
			SecretKey:  os.Getenv("JWT_SECRET_KEY"),
		},
		Complex: ComplexConfig{},
	}
}
