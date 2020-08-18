package main

import (
	"ami"
	"net/http"
)

func main() {
	r := ami.Default()
	r.GET("/", func(c *ami.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})

	r.GET("/panic", func(c *ami.Context) {
		names := []string{"arcades"}
		c.String(http.StatusOK, names[1000])
	})

	_ = r.Run(":9999")
}
