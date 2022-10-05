package database

import (
	"UrlShortner/model"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client = NewClient()

// docker run -d -p 6379:6379 redislabs/redismod
func NewClient() *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		},
	)
}

func GetOriginalURL(ShortURL string) (string, error) {
	LongURL, err := client.Get(ctx, ShortURL).Result()
	if err != nil {
		return "", err
	}
	return LongURL, nil
}

func SetURLs(URLs *model.URL) error {
	err := client.Set(ctx, URLs.ShortURL, URLs.LongURL, 0)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
