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

// SignUpHandler 用户注册
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
		passwdSha := encPasswd(passwd)
		if ok := db.UserSignUp(username, passwdSha); ok {
			io.WriteString(w, "SUCCESS")
		} else {
			io.WriteString(w, "FAILED")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// SignInHandler 用户登陆
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if html, err := ioutil.ReadFile("./static/view/signin.html"); err == nil {
			w.Write(html)
		} else {
			w.Write([]byte(err.Error()))
		}
		return
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		passwd := r.Form.Get("password")
		passwdSha := encPasswd(passwd)
		if ok := db.UserSignIn(username, passwdSha); ok {
			io.WriteString(w, "ok")
		} else {
			io.WriteString(w, "Failed to login")
			// log.Println("Failed to login")
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func encPasswd(passwd string) string {
	return util.Sha1([]byte(passwd + pwdSalt))
}
