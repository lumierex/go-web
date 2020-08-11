package main

import (
	"ami"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := ami.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.path = %q \n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	err := r.Run(":8889")
	if err != nil {
		log.Fatal(err)
	}
}
