package main

import (
	"Ozon_Fintech/app/model"
	"Ozon_Fintech/app/service"
	"fmt"
	"log"
	"net/http"
)

func short_url(writter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		LongURL := request.URL.Query().Get("URL")
		URLs := &model.URL{LongURL: LongURL}
		err := service.GenerateShortURL(URLs)
		if err != nil {
			return
		}
		URLs.ShortURL = "http://" + request.Host + "/" + URLs.ShortURL
		writter.Header().Set("Content-Type", "value")
		writter.WriteHeader(http.StatusOK)
		writter.Write([]byte(URLs.ShortURL))
	} else {
		writter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func reddirect(writter http.ResponseWriter, request *http.Request) {
	url := request.URL
	res, err := service.GetOriginalURL(url.Path[1:])
	if err != nil {
		return
	}
	LongURL := "http://" + res
	http.Redirect(writter, request, LongURL, http.StatusTemporaryRedirect)
}

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "3525"
	dbname   = "pos"
)

func main() {
	http.HandleFunc("/", reddirect)
	http.HandleFunc("/short-url", short_url)
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
