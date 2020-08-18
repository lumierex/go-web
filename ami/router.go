package ami

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(path string) []string {
	sp := strings.Split(path, "/")
	parts := make([]string, 0)
	for _, part := range sp {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	fmt.Printf("parsePattern : %v \n", parts)
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {

	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	fmt.Printf("getRoute n: %v", n)

	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {

			// /hello/*test ==> 只是接收 *test
			if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}

			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
		}

		fmt.Printf("getRoute params %v", params)
		return n, params
	}

	return nil, nil

}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		// key规定method + pattern
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
		//r.handlers[key](c)
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
		})
	}
	// TODO 此处的next?
	c.Next()
}
