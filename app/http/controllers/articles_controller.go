package controllers

import (
	"database/sql"
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"html/template"
	"net/http"
)

// ArticlesControllers 文章相关页面
type ArticlesControllers struct {
}

// Show 文章详情页面
func (*ArticlesControllers) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章的数据
	article, err := getArticlesByID(id)

	// 3. 如果出现了错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		}
	} else {
		// 4. 读取成功
		tmpl, err := template.New("show.html").Funcs(template.FuncMap{
			"RouteName2URL": route.RouteName2URL,
			"Int64ToString": types.Int64ToString,
		}).ParseFiles("resources/views/articles/show.html")
		logger.LogError(err)
		tmpl.Execute(w, article)
	}
}
