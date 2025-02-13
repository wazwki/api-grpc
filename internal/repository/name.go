package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type NameRepository struct {
	DataBase *pgxpool.Pool
	Cache    *redis.Client
}

func NewNameRepository(db *pgxpool.Pool, cache *redis.Client) NameRepositoryInterface {
	return &NameRepository{DataBase: db, Cache: cache}
}
