package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	Title string
	Body  string
	models.BaseModel
}

// 	Link 方法用来生成文章连接
// func (a Article) Link() string {
// 	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
// }

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
