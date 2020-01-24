package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"netdisk/db"
	"netdisk/util"
	"strconv"
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
		token := genToken(username)
		if ok := db.UpdateToken(username, token); !ok {
			io.WriteString(w, "Failed to updsate token")
			return
		}
		// 3. 登陆成功后重定向到首页
		response := ResponseJSON(map[string]interface{}{
			"Token":    token,
			"Username": username,
			"Location": "/home",
		})

		io.WriteString(w, response)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// UserInfoHandler 用户信息接口
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		// 1. 解析参数
		username := r.Form.Get("username")
		token := r.Form.Get("token")

		// 2. 验证 token
		if isInvalid := IsTokenValid(username, token); !isInvalid {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// 3. 查询用户信息
		userInfo, err := db.GetUserInfo(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 4. 组装并响应用户数据
		response := util.RespMsg{
			Code: 0,
			Msg:  "OK",
			Data: userInfo,
		}
		w.Write(response.JSONBytes())
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// IsTokenValid 验证 token 时效性
func IsTokenValid(username, token string) bool {
	// 1579840405   704610000
	tsToken, _ := strconv.Atoi(token[32:])
	tsCur := time.Now().UnixNano() / int64(math.Pow10(11))
	// ts - (unit convert) -> ts(s) --> ts(hour)
	// 如果 token > 一小时，token 失效
	if diffValue := tsCur - int64(tsToken); (diffValue * 100 / 3600) > 1 {
		return false
	}
	// 验证 token 的正确性
	if tokenFromDB, err := db.GetToken(username); err != nil || token != tokenFromDB {
		return false
	}
	return true
}

// GenToken 生成 token
func genToken(username string) string {
	// md5(username + timestamp + token_salt) + timestamp[:8]
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())
	token := util.MD5([]byte(username+timestamp+"_tokensalt")) + timestamp[:8]
	return token
}

func encPasswd(passwd string) string {
	return util.Sha1([]byte(passwd + pwdSalt))
}
