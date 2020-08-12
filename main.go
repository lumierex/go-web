package main

import (
	"ami"
	"log"
	"net/http"
)

func main() {
	r := ami.New()
	//r.GET("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "URL.path = %q \n", r.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	for k, v := range r.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	r.GET("/", func(c *ami.Context) {
		c.HTML(http.StatusOK, `<div style="width:100vw;height:100vh;background:yellow">hello world</div>`)
	})
	r.GET("/string", func(c *ami.Context) {
		c.String(http.StatusOK, "hello string %s", "酸爽")
	})
	r.POST("/login", func(c *ami.Context) {
		c.JSON(http.StatusOK, ami.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/info", func(c *ami.Context) {
		c.JSON(http.StatusOK, ami.H{
			"username": c.Query("username"),
			"password": c.Query("password"),
		})
	})
	r.GET("/data", func(c *ami.Context) {
		c.Data(http.StatusOK, []byte("hi"))
	})
	err := r.Run(":8889")
	if err != nil {
		log.Fatal(err)
	}
}
