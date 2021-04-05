package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (this *Server) Listen() error {
	http.Handle("/", this.router)
	err := http.ListenAndServe(this.port, nil)
	return err
}

func (this *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := this.router.rules[path]
	if !exists {
		this.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	this.router.rules[path][method] = handler
}

func (this *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
