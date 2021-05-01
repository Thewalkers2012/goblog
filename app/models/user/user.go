package user

import "goblog/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string
	Email    string
	Password string
}
