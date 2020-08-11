package ami

import (
	"fmt"
	"log"
	"net/http"
)

// 用户实现的函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP实现结构体
type Engine struct {
	router map[string]HandlerFunc
}

// 引擎构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// 添加GET路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 添加POST路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string)(err error) {
	return http.ListenAndServe(addr, engine);
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s\n", r.URL)
	}

}