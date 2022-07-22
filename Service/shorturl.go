package Service

import "math/rand"

type original_short_url struct {
	original_short string `json:"Original_Short"`
	short_url      string `json:"Short_Url"`
}

func generte_short_url(url string) string {
	letters := []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789_")
	result := make([]rune, 10)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
