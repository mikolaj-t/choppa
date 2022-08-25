package main

import (
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
	"os"
	"strconv"
)

var redisPwd = os.Getenv("REDIS_PASSWORD")
var db = redisdb{address: "redis:6379", password: redisPwd}

type redisdb struct {
	address  string
	password string
	redis    *redis.Client
}

func (db *redisdb) connect() {
	db.redis = redis.NewClient(&redis.Options{
		Addr:       db.address,
		Password:   db.password,
		DB:         0, // default
		MaxRetries: 10,
	})

	_, err := db.redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func (db *redisdb) updateChop(path string, platform int, url string) error {
	return db.redis.Set(context.Background(), getKeyName(path, platform), url, 0).Err()
}

func (db *redisdb) getChop(path string, platform int) (string, error) {
	url, err := db.redis.Get(context.Background(), getKeyName(path, platform)).Result()
	return url, err
}

func getKeyName(path string, platform int) string {
	return path + ":" + strconv.Itoa(platform)
}
