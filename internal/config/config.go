package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host               string
	Port               string
	GRPCPort           string
	DBdsn              string
	LogLevel           string
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
	AccessTokenTTL     int
	RefreshTokenTTL    int
	RedisHost          string
	RedisPort          string
	RedisPassword      string
	RedisDBNumber      int
	Debug              bool
}

func LoadFromEnv() (*Config, error) {
	attl, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		return nil, err
	}
	rttl, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		return nil, err
	}
	dbnum, err := strconv.Atoi(os.Getenv("REDIS_NUMBER"))
	if err != nil {
		return nil, err
	}
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		GRPCPort: os.Getenv("GRPC_PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
		DBdsn: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")),
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		RedisDBNumber:      dbnum,
		AccessTokenSecret:  []byte(os.Getenv("JWT_ACCESS_SECRET")),
		RefreshTokenSecret: []byte(os.Getenv("JWT_REFRESH_SECRET")),
		AccessTokenTTL:     attl,
		RefreshTokenTTL:    rttl,
		Debug:              debug,
	}

	return cfg, nil
}
