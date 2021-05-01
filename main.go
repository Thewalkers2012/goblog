package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	// 路由的初始化
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	// 通过命名路由获取 URL 实例
	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
