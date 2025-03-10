package models

// 定义音频书籍模型
type AudioBook struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
}
