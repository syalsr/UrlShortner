package database

import (
	"Ozon_Fintech/app/model"
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var client = NewClient()

// docker run -d -p 6379:6379 redislabs/redismod
func NewClient() *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
}

func GetOriginalURL(ShortURL string) (string, error) {
	LongURL, err := client.Get(Ctx, ShortURL).Result()
	if err != nil {
		return "", err
	}
	return LongURL, nil
}

func SetURLs(URLs *model.URL) error {
	err := client.Set(Ctx, URLs.ShortURL, URLs.LongURL, 0)
	if err != nil {
		return err.Err()
	}
	return nil
}
