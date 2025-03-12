package models

// 定义有声内容类别，如有声书、广播剧、评书、相声、音乐
type Class struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	NameId int    `gorm:"not null;unique"`
	Name   string `gorm:"not null"`
}
