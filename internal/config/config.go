package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

const (
	defaultPort = "8000"
)

// Config is the central setting.
type Config struct {
	Port                 string
	RefreshSigningString string
	AccessSigningString  string
}

func GetConfig() *Config {

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}

	conf := &Config{
		Port:                 port,
		AccessSigningString:  os.Getenv("SIGNING_STRING"),
		RefreshSigningString: os.Getenv("REFRESH_SIGNING_STRING"),
	}

	return conf
}
