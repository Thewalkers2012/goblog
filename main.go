package main

import (
	"fmt"
	"goblog/bootstrap"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var router *mux.Router

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径的斜杠
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 路由的初始化
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	router.Use(forceHTMLMiddleware) // 使用Gorillia Mux 中的use方法来加载中间件
	// 通过命名路由获取 URL 实例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL: ", articleURL)
	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
