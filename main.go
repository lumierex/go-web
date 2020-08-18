package main

import (
	"ami"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

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

	//r := ami.New()
	//r.GET("/index", func(c *ami.Context) {
	//	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	//})
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", func(c *ami.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//
	//	v1.GET("/hello", func(c *ami.Context) {
	//		// expect /hello?name=amiktutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//v2 := r.Group("/v2")
	//{
	//	v2.GET("/hello/:name", func(c *ami.Context) {
	//		// expect /hello/amiktutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//	v2.POST("/login", func(c *ami.Context) {
	//		c.JSON(http.StatusOK, ami.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//
	//}
	//
	//_ = r.Run(":9999")

	// 中间件
	r := ami.New()
	r.Use(ami.Logger())
	//
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	s1 := &student{
		Name: "min",
		Age:  12,
	}

	s2 := &student{
		Name: "wei",
		Age:  11,
	}
	r.GET("/", func(c *ami.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *ami.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", ami.H{
			"title":  "ami",
			"stuArr": [2]*student{s1, s2},
		})
	})

	r.GET("/date", func(c *ami.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", ami.H{
			"title": "ami",
			"now":   time.Date(2020, 8, 17, 0, 0, 0, 0, time.UTC),
			//"now", time.Date(2020, 08, 17, 0, 0, 0, 0, time.UTC):
		})
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForR1())
	{
		v2.GET("/hello/:name", func(c *ami.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	// 静态资源文件
	// 将磁盘上的文件root映射到relativePath上
	//r.Static("/asset", "D:\\code\\react\\jdme-otp" )
	//r.Static("/asset", "./" )
	//r.Static("/asset", "D:/code/react/jdme-otp")
	_ = r.Run(":9999")

}

func onlyForR1() ami.HandlerFunc {
	return func(c *ami.Context) {
		// Start timer
		t := time.Now()
		//// if a server error occurred
		//log.Fatal(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group r2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
