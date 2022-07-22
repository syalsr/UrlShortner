package main

import (
	"fmt"
	"log"
	"net/http"
)

func url(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)

		return
	}
	switch r.Method {
	case "Get":
		http.ServeFile(w, r, "forms.html")
	case "Post":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}
		fmt.Fprintf(w, "Post from webstie r.postfrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Only get and post")
	}

}

func short_url(writter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		url = request.URL.Query().Get("URL")
		//host := request.Host
		shortURL :=
			fmt.Println(1)
	} else {
		writter.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	fmt.Println("Application for short URL")
	http.HandleFunc("/short_url", short_url)
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
