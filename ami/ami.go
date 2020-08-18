package ami

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
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

		// 模板渲染
		htmlTemplates *template.Template
		funcMap       template.FuncMap
	}
)

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

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

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 静态文件解析
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {

	absolutePath := path.Join(group.prefix, relativePath)
	fmt.Printf("create Static Handler %s, group.prefix%s\n", absolutePath, group.prefix)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
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

// 模板加载进内存
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// 所有模板渲染函数
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

// 生成上下文
// 把上下文传输给路由管理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc

	// 每个group都有对应的中间件
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}
