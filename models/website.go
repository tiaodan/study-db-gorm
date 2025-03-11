package models

// 定义网站模型
type Website struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	NameId int    `gorm:"not null"`
	Name   string `gorm:"not null"`
	URL    string `gorm:"not null"`
}
