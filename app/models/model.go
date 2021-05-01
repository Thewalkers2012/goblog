package models

import (
	"goblog/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID       uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreateAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
