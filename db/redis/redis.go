package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Config(dbhost, dbport, dbpassword string, dbnumber int) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", dbhost, dbport),
		Password: dbpassword,
		DB:       dbnumber,
	}), nil
}
