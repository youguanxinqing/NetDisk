package handler

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"netdisk/db"
	"netdisk/util"
)

const (
	pwdSalt = "*#890"
)

// SignUpHandler ...
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if html, err := ioutil.ReadFile("./static/view/signup.html"); err == nil {
			w.Write(html)
		} else {
			w.Write([]byte(err.Error()))
		}
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		passwd := r.Form.Get("password")
		log.Println(username + ":" + passwd)
		// encode passwd
		passwdSha := util.Sha1([]byte(passwd + pwdSalt))
		if ok := db.UserSignUp(username, passwdSha); ok {
			io.WriteString(w, "SUCCESS")
		} else {
			io.WriteString(w, "FAILED")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
