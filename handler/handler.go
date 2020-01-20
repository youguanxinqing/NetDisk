package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"netdisk/meta"
	"netdisk/util"
	"os"
	"time"
)

// UploadHandler 上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Http Get
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
		// Http Post
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Sprintf("failed get data : %v", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2018-01-02 12:20"),
		}
		// 创建文件
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Sprintf("failed create new file : %v", err.Error())
			return
		}
		defer newFile.Close()
		// 写入文件
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Sprintf("failed store file : %v", err.Error())
			return
		}
		// 读指针移动到文件初始位置
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "finished upload")
}

// GetFileMetaHandler 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	// 提取 url params
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(filehash)
	// struct -> json
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// FileDownloadHandler 文件下载
func FileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(filehash)

	fd, err := os.Open(fileMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(fd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 为下载文件设置响应头信息
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.FileName+"\"")
	w.Write(data)
}

// FileMetaUpdateHandler 修改文件元信息(文件重命名)
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	filehash := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 更新 file meta
	fileMeta := meta.GetFileMeta(filehash)
	fileMeta.FileName = newFileName
	meta.UpdateFileMeta(fileMeta)
	// struct -> json
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// FileDeleteHandler 删除文件..
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	// ...
	fileMeta := meta.GetFileMeta(filehash)
	// 删除文件
	os.Remove(fileMeta.Location)
	// 清除文件元信息
	meta.RemoveFileMeta(filehash)

	w.WriteHeader(http.StatusOK)
}
