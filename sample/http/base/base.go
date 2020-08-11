package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {}

// ServeHTTP  ListenAndServe的回调函数
//
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k,v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s", req.URL)
	}
}

func main()  {
	//http.HandleFunc("/", IndexHandler)
	//http.HandleFunc("/hello", helloHandler)
	//
	//http.ListenAndServe(":8888", nil)
	engine := &Engine{}
	log.Fatal(http.ListenAndServe(":8888", engine))
}

func IndexHandler(w http.ResponseWriter, req * http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req * http.Request)  {
	for k,v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}