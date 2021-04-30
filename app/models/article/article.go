package article

import (
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID    int64 // 注意类型，FormarInt 的转化类型为 int64 这里也需要对应，否则不能转换
	Title string
	Body  string
}

// 	Link 方法用来生成文章连接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
