package conf

import (
	"log"
	"os"
	"strconv"

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

	userMaxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	userMinConn, _ := strconv.Atoi(os.Getenv("DB_MIN_CONN"))

	return &Config{
		User: UserConfig{
			ServerPort:    os.Getenv("SERVER_PORT"),
			SecretKey:     os.Getenv("JWT_SECRET_KEY"),
			DebugPort:     os.Getenv("DEBUG_PORT"),
			DbDsn:         os.Getenv("DB_DSN"),
			DbMaxOpenConn: userMaxConn,
			DbMinOpenConn: userMinConn,
		},
		Complex: ComplexConfig{},
	}
}
