package main

import (
	"ami"
	"net/http"
)

func main() {
	//r := ami.New()
	//r.GET("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "URL.path = %q \n", r.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	for k, v := range r.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//r.GET("/", func(c *ami.Context) {
	//	c.HTML(http.StatusOK, `<div style="width:100vw;height:100vh;background:yellow">hello world</div>`)
	//})
	//r.GET("/string", func(c *ami.Context) {
	//	c.String(http.StatusOK, "hello string %s", "酸爽")
	//})
	//r.POST("/login", func(c *ami.Context) {
	//	c.JSON(http.StatusOK, ami.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	//
	//r.GET("/info", func(c *ami.Context) {
	//	c.JSON(http.StatusOK, ami.H{
	//		"username": c.Query("username"),
	//		"password": c.Query("password"),
	//	})
	//})
	//r.GET("/data", func(c *ami.Context) {
	//	c.Data(http.StatusOK, []byte("hi"))
	//})
	//err := r.Run(":8889")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//r := ami.New()
	//
	//r.GET("/", func(c *ami.Context) {
	//	c.HTML(http.StatusOK, `<h1 style="color:red">hello world</h1>`)
	//})
	//
	//r.GET("/hello", func(c *ami.Context) {
	//	c.String(http.StatusOK, "hello %s, you are at %s \n", c.Path)
	//})
	//
	//r.GET("/hello/:name", func(c *ami.Context) {
	//	c.String(http.StatusOK, "hello %s you are at %s\n", c.Params["name"], c.Path)
	//})
	//
	//r.GET("/assets/*filepath", func(c *ami.Context) {
	//	c.JSON(http.StatusOK, ami.H{
	//		"filepath": c.Param("filepath"),
	//	})
	//})
	//_ = r.Run(":9999")

	r := ami.New()
	r.GET("/index", func(c *ami.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ami.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *ami.Context) {
			// expect /hello?name=amiktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ami.Context) {
			// expect /hello/amiktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ami.Context) {
			c.JSON(http.StatusOK, ami.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	_ = r.Run(":9999")

}
