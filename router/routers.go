package router

import "net/http"

type middlewareFunc func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	middlewares []middlewareFunc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(mw middlewareFunc) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) Add(pattern string, handler http.HandlerFunc) {
	for i := len(r.middlewares); i >= 0; i-- {
		handler = (r.middlewares[i])(handler)
	}
	http.HandleFunc(pattern, handler)
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, nil)
}
