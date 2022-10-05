package service

import (
	"Ozon_Fintech/app/database"
	"Ozon_Fintech/app/model"
	"math/rand"
)

func GenerateShortURL(URLs *model.URL) error {
	letters := []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789")
	result := make([]rune, 10)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	URLs.ShortURL = string(result)
	err := database.SetURLs(URLs)
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalURL(ShortURL string) (string, error) {
	LongURL, err := database.GetOriginalURL(ShortURL)
	return LongURL, err
}
