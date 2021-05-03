package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 路由的初始化
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	// 通过命名路由获取 URL 实例
	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
