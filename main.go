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
		http.HandleFunc("/home", handler.HomeHandler)
	}

	{ // core
		http.HandleFunc("/file/upload", handler.UploadHandler)
		http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
		http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
		http.HandleFunc("/file/download", handler.FileDownloadHandler)
		http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
		http.HandleFunc("/file/delete", handler.FileDeleteHandler)
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
