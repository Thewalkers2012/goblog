package models

import (
	"goblog/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID       uint64
	CreateAt time.Time
	UpdataAt time.Time
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
