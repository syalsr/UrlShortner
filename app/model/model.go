package model

import "encoding/json"

type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func ConvertToJson(URLs *URL) ([]byte, error) {
	data, err := json.MarshalIndent(URLs, "", "\t")
	if err != nil {
		return nil, err
	}
	return data, nil
}
