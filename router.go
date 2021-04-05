package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (this *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	var handler http.HandlerFunc
	var existsMethod bool
	_, existsPath := this.rules[path]

	if existsPath {
		handler, existsMethod = this.rules[path][method]
	}
	return handler, existsPath, existsMethod
}

func (this *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, existsPath, existsMethod := this.FindHandler(request.Method, request.URL.Path)
	if existsMethod {
		handler(w, request)
	} else if !existsPath {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Route Not Found")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
	}
}
