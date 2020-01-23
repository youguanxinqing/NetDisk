package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"netdisk/db"
	"netdisk/util"
	"time"
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
		// 1. 账号密码校验
		if ok := db.UserSignIn(username, passwdSha); !ok {
			io.WriteString(w, "Failed to login")
			return
		}
		// 2. 生成 token
		token := GenToken(username)
		if ok := db.UpdateToken(username, token); !ok {
			io.WriteString(w, "Failed to updsate token")
			return
		}
		// 3. 登陆成功后重定向到首页
		response := ResponseJSON(map[string]interface{}{
			"Toke":     token,
			"Username": username,
			"Location": "/home",
		})
		io.WriteString(w, response)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func encPasswd(passwd string) string {
	return util.Sha1([]byte(passwd + pwdSalt))
}

// GenToken 生成 token
func GenToken(username string) string {
	// md5(username + timestamp + token_salt) + timestamp[:8]
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())
	token := util.MD5([]byte(username+timestamp+"_tokensalt")) + timestamp[:8]
	return token
}
