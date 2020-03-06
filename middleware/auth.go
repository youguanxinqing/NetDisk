package middleware

import (
	"net/http"
	"netdisk/handler"
	"netdisk/util"
)

func AuthMiddleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")

		if len(username) < 3 || !handler.IsTokenValid(username, token) {
			response := util.NewRespMsg(
				http.StatusMovedPermanently, "token is expired", map[string]string{
					"Location": "/user/signin",
				})
			w.Write(response.JSONBytes())
			return
		}

		h(w, r)
	}
}
