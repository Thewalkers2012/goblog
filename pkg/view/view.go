package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render 渲染视图
func Render(w io.Writer, name string, data interface{}) {
	// 1. 设置模版相对路径
	viewDir := "resources/views/"

	// 2. 语法糖，将 articles.show 更改为 articles/show
	name = strings.Replace(name, ".", "/", -1) // 最后一个参数 -1 代表替换所有

	// 3. 所有布局模版文件 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.html")
	logger.LogError(err)

	// 4. 在 Slice 里新增我们的目标文件
	newFiles := append(files, viewDir+name+".html")

	// 5. 解析所有模版文件
	tmpl, err := template.New(name + ".html").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)

	// 6. 渲染模版
	tmpl.ExecuteTemplate(w, "app", data)
}
