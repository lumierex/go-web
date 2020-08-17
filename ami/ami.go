package ami

import (
	"fmt"
	"log"
	"net/http"
)

// 用户实现的函数
//type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(c *Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		// 每个router-group共享一个engine的实例
		engine *Engine
	}

	// ServeHTTP实现结构体

	Engine struct {
		// 嵌套属性 engine拥有RouterGroup属性
		// 所以engine就有 group的Group属性
		//
		*RouterGroup
		router *router
		// Engine拥有所有的route-group
		groups []*RouterGroup
	}
)

// 引擎构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}

	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
	//return &Engine{router: newRouter()}
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	fmt.Printf("group prefix: %s, group: %v\n", group.prefix, group)
	newGroup := &RouterGroup{
		// 嵌套group
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// comp -> pattern
func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//engine.
	engine.router.addRoute(method, pattern, handler)
}

// 添加GET路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 添加POST路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 生成上下文
// 把上下文传输给路由管理

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)

}
