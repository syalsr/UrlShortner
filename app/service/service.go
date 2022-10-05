package service

import "math/rand"

func GenerateShortURL(url string) string {
	letters := []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789")
	result := make([]rune, 10)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func GetOriginalURL(ShortURL string) string {
	return ""
}
