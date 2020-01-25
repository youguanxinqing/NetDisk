package handler

import (
	"net/http"
)

// HTTPInterceptor 拦截器(中间件)
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	// http.HandlerFunc(func(w http.ResponseWriter, r *httpRequest)) 类型转换
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")

		if len(username) < 3 || !IsTokenValid(username, token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		h(w, r)
	})
}
