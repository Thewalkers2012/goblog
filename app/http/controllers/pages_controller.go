package controllers

import (
	"fmt"
	"net/http"
)

// PagesController  处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！ </h1>")
}

// AboutHandler 关于我们的页面
func (*PagesController) AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系"+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

// NotFoundHandler 404 页面的设置
func (*PagesController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

/*
	目录结构的设置
	app 目录存放业务逻辑相关的代码，子目录还会有 models 等目录，相对于 pkg 存放底层代码而言，不同的项目，pkg 目录存放的差异不大
	app/http 目录用来存放 controllers, requests (表单验证)， middlewares 等跟处理 HTTP 请求相关的代码
	app/http/controllers 目录存放着处理 HTTP 请求的控制器，每个控制器文件里包含多个方法，这些方法是最终处理用户请求病返回相应的地方
*/
