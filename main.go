package main

import (
	"fmt"
	"net/http"

	"netdisk/handler"
)

func main() {
	// routes
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.FileDownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	// listen
	addr := ":8080"
	fmt.Println("start service base on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
