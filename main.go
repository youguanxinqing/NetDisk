package main

import (
	"fmt"
	"net/http"

	"netdisk/handler"
	_ "netdisk/settings"
)

func main() {
	// routes
	{ // base
		http.HandleFunc("/static/", handler.StaticResource)
		// FIXME: home 页面不能做 auth, 否则不能登陆成功
		http.HandleFunc("/home", handler.HomeHandler)
	}

	{ // core
		http.HandleFunc("/file/upload", handler.HTTPInterceptor(handler.UploadHandler))
		http.HandleFunc("/file/fast/upload", handler.HTTPInterceptor(handler.TryFastUploadHander))
		http.HandleFunc("/file/upload/suc", handler.HTTPInterceptor(handler.UploadSucHandler))
		http.HandleFunc("/file/meta", handler.HTTPInterceptor(handler.GetFileMetaHandler))
		http.HandleFunc("/file/download", handler.HTTPInterceptor(handler.FileDownloadHandler))
		http.HandleFunc("/file/update", handler.HTTPInterceptor(handler.FileMetaUpdateHandler))
		http.HandleFunc("/file/delete", handler.HTTPInterceptor(handler.FileDeleteHandler))
		http.HandleFunc("/file/query", handler.HTTPInterceptor(handler.QueryUserFileMetasHandler))
	}

	{ // user
		http.HandleFunc("/user/signup", handler.SignUpHandler)
		http.HandleFunc("/user/signin", handler.SignInHandler)
		http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))
	}

	// listen
	addr := ":8080"
	fmt.Println("start service base on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
