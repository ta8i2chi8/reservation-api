package handler

import (
	"net/http"
	"strings"
)

type Route struct {
	Path        string
	Method      string
	Handler     http.HandlerFunc
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Path:        path,
		Method:      method,
		Handler:     handler,
		Middlewares: middlewares,
	})
}

func (r *Router) GET(path string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	r.AddRoute("GET", path, handler, middlewares...)
}

func (r *Router) POST(path string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	r.AddRoute("POST", path, handler, middlewares...)
}

func (r *Router) PUT(path string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	r.AddRoute("PUT", path, handler, middlewares...)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	r.AddRoute("DELETE", path, handler, middlewares...)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if r.matchRoute(route, req) {
			handler := route.Handler

			for _, middleware := range route.Middlewares {
				handler = middleware(handler)
			}

			handler.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func (r *Router) matchRoute(route Route, req *http.Request) bool {
	if route.Method != req.Method {
		return false
	}

	return r.matchPath(route.Path, req.URL.Path)
}

func (r *Router) matchPath(routePath, requestPath string) bool {
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")
	requestParts := strings.Split(strings.Trim(requestPath, "/"), "/")

	if len(routeParts) != len(requestParts) {
		return false
	}

	for i := 0; i < len(routeParts); i++ {
		if len(routeParts[i]) > 0 && routeParts[i][0] == ':' {
			continue
		}
		if routeParts[i] != requestParts[i] {
			return false
		}
	}

	return true
}
