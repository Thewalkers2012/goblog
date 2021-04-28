package article

import (
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID    int
	Title string
	Body  string
}

// 	Link 方法用来生成文章连接
func (a Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
