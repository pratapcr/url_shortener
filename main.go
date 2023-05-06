package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/server/helper"
	"example.com/server/link"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing request body", http.StatusInternalServerError)
				return
			}
			formData := r.Form
			fmt.Println(formData.Get("url"))
			longUrl := formData.Get("url")
			shortUrl := "http://localhost:1234/" + link.RandomString(10)
			fmt.Println(shortUrl)
			helper.Insert(longUrl, shortUrl)
			formData.Set("shortUrl", shortUrl)
			fmt.Fprintf(w, "response: %v\n", formData)
		}

		if r.Method == http.MethodGet {
			if r.URL.Path == "/favicon.ico" {
				return
			}
			fmt.Println(r.URL.Path)
			helper.GetLongUrl("http://localhost:1234" + r.URL.Path)
			// http.Redirect()
		}
	})

	helper.Init()
	log.Println("Starting the server on port:", 1234)
	if err := http.ListenAndServe(": 1234", nil); err != nil {
		log.Println("Error while starting the server :", err)
		panic(err)
	}
}
