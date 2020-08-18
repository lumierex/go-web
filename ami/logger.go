package ami

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// 定时
		t := time.Now()
		// 执行请求
		c.Next()
		// 打印请求耗时
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
