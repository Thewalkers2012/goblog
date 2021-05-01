package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 1. 表单验证
	// 2. 验证通过 —— 入库，并跳转到首页
	// 3. 表单不通过 —— 重新显示表单
}
