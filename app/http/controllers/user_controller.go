package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	BaseController
}

// Show 用户个人页面
func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {

	// 1. 获取用户 URL
	id := route.GetRouteVariable("id", r)

	// 2. 读取相应的文章数据
	_user, err := user.Get(id)

	// 3. 如果出现错误
	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		// 4. 读取成功
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
