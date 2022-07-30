package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func generte_short_url(url string) string {
	letters := []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789_")
	result := make([]rune, 10)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func add_to_file(originalurl, shorturl string) {
	file_name := "database"
	result := originalurl + "|" + shorturl + "\n"
	if _, err := os.Stat(file_name); errors.Is(err, os.ErrNotExist) {
		os.Create(file_name)
		os.WriteFile(file_name, []byte(result), 0644)
	} else {
		input, err := os.ReadFile(file_name)
		if err != nil {
			return
		}
		lines := strings.Split(string(input), "\n")

		fileIsChange := false
		for idx, item := range lines {
			if strings.Contains(item, shorturl) {
				fileIsChange = true
				lines[idx] = result
			}
		}
		if fileIsChange {
			output := strings.Join(lines, "\n")
			os.WriteFile(file_name, []byte(output), 0644)
		} else {
			output := strings.Join(lines, "\n")
			output += result
			os.WriteFile(file_name, []byte(output), 0644)
		}

	}
}

func short_url(writter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		url := request.URL.Query().Get("URL")
		host := request.Host
		shortURL := generte_short_url(url)
		add_to_file(url, shortURL)
		res_str := "http://" + host + "/" + shortURL
		writter.Header().Set("Content-Type", "value")
		writter.WriteHeader(http.StatusOK)
		writter.Write([]byte(res_str))
	} else {
		writter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func get_original_url(short_url string) string {
	str, _ := os.ReadFile("database")
	lines := strings.Split(string(str), "\n")
	for _, item := range lines {
		if strings.Contains(item, short_url) {
			urls := strings.Split(item, "|")
			return urls[0]
		}
	}
	return "3"
}

func reddirect(writter http.ResponseWriter, request *http.Request) {
	url := request.URL
	res := get_original_url(url.Path[1:])
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

func postgre() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}

func main() {
	postgre()
	fmt.Println("Application for short URL")
	http.HandleFunc("/", reddirect)
	http.HandleFunc("/short-url", short_url)
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
