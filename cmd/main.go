package main

import (
	"Ozon_Fintech/app/service"
	"fmt"
	"log"
	"net/http"
)

func short_url(writter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		url := request.URL.Query().Get("URL")
		shortURL := service.GenerateShortURL(url)
		res_str := "http://" + request.Host + "/" + shortURL
		writter.Header().Set("Content-Type", "value")
		writter.WriteHeader(http.StatusOK)
		writter.Write([]byte(res_str))
	} else {
		writter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func reddirect(writter http.ResponseWriter, request *http.Request) {
	url := request.URL
	res := service.GetOriginalURL(url.Path[1:])
	newres := "http://" + res
	http.Redirect(writter, request, newres, http.StatusTemporaryRedirect)
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
