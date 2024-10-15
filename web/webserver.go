package web

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

// RouterGroup represents a group of routes with a common prefix and middleware.
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	router      *router
}

// Server represents the main web server that manages RouterGroups and routes.
type Server struct {
	*RouterGroup
}

func New() *Server {
	server := &Server{}
	server.RouterGroup = &RouterGroup{router: newRouter()}
	return server
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		router: group.router,
	}
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.router.addRoute(method, pattern, handler)
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 替代switch-case
	// if handler, ok := server.router[r.Method+"-"+r.URL.Path]; ok {
	// 	handler(w, r)
	// } else {
	// 	w.WriteHeader(http.StatusNotFound)
	// }
	c := newContext(w, r)
	server.router.handle(c)
}

func (server *Server) Run(addr string) (err error) {
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, server)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Use registers middleware for the RouterGroup.
func (group *RouterGroup) Use(middleware HandlerFunc) {
	group.middlewares = append(group.middlewares, middleware)
}
