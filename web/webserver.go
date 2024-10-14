package web

import "net/http"

type HandlerFunc func(*Context)

type Server struct {
	router *router
}

func New() *Server {
	return &Server{router: newRouter()}
}

func (server *Server) addRoute(method string, pattern string, handler HandlerFunc) {
	server.router.addRoute(method, pattern, handler)
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
	return http.ListenAndServe(addr, server)
}

func (server *Server) GET(pattern string, handler HandlerFunc) {
	server.addRoute("GET", pattern, handler)
}

func (server *Server) POST(pattern string, handler HandlerFunc) {
	server.addRoute("POST", pattern, handler)
}
