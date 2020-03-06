package router

import (
	"netdisk/handler"
	"netdisk/middleware"
)

func DefaultRouter() *Router {
	r := NewRouter()

	// 挂载中间件
	r.Use(middleware.AuthMiddleWare)

	// 注册路由
	register(r)

	return r
}

func register(r *Router) {
	// routes
	{ // base
		r.Add("/static/", handler.StaticResource)
		// FIXME: home 页面不能做 auth, 否则不能登陆成功
		r.Add("/home", handler.HomeHandler)
	}

	{ // core
		r.Add("/file/upload", handler.UploadHandler)
		r.Add("/file/fast/upload", handler.TryFastUploadHander)
		r.Add("/file/upload/suc", handler.UploadSucHandler)
		r.Add("/file/meta", handler.GetFileMetaHandler)
		r.Add("/file/download", handler.FileDownloadHandler)
		r.Add("/file/update", handler.FileMetaUpdateHandler)
		r.Add("/file/delete", handler.FileDeleteHandler)
		r.Add("/file/query", handler.QueryUserFileMetasHandler)
	}

	{ // user
		r.Add("/user/signup", handler.SignUpHandler)
		r.Add("/user/signin", handler.SignInHandler)
		r.Add("/user/info", handler.UserInfoHandler)
	}
}
