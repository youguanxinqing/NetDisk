package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// StaticResource 返回静态资源文件
func StaticResource(w http.ResponseWriter, r *http.Request) {
	comps := strings.Split(r.URL.Path, "/")
	// 拼接文件路径
	resourcePath := strings.Join(comps, string(os.PathSeparator))
	data, err := ioutil.ReadFile("." + resourcePath)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	w.Write(data)
}
