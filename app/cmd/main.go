package main

import (
	"UrlShortner/model"
	"UrlShortner/service"
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

		data, err := model.ConvertToJson(URLs)
		if err != nil {
			return
		}
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusOK)
		writter.Write(data)
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

func main() {
	http.HandleFunc("/", reddirect)
	http.HandleFunc("/short-url", short_url)
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
