package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
	models.BaseModel
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
