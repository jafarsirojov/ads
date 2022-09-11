package config

import (
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"io"
	"os"
)

type Config struct {
	Port    string
	Version string

	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var Module = fx.Options(fx.Provide(ProvideConfig))

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ProvideConfig() *Config {
	return &Config{
		DBHost:     getEnv("DATABASE_HOST", "localhost"),
		DBUser:     getEnv("DATABASE_USER", "postgres"),
		DBPassword: getEnv("DATABASE_PASSWORD", "pass"),
		DBPort:     getEnv("DATABASE_PORT", "5432"),
		DBName:     getEnv("DATABASE_NAME", "adsdb"),

		Port:    getEnv("PORT", ":1111"),
		Version: getEnv("VERSION", "v1"),
	}
}

func (c *Config) Dump(w io.Writer) error {
	fmt.Fprint(w, "config: ")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(c)
}
