package db

import (
	"log"
	"study-db-gorm/models"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once // 使用 sync.Once 确保单例

// 初始化数据库连接
func InitDB() {
	once.Do(func() { // 使用 sync.Once 确保只执行一次
		dsn := "root:password@tcp(127.0.0.1:3306)/audio?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("数据库连接失败:", err)
			panic(err)
		}
		log.Println("数据库连接成功")
	})
}

// 创建记录
func CreateAudioBook(book *models.AudioBook) {
	result := DB.Create(book)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
	} else {
		log.Println("创建成功:", book)
	}
}

// 查询记录
func GetAudioBookByID(id uint) *models.AudioBook {
	var book models.AudioBook
	result := DB.First(&book, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", book)
	return &book
}

// 更新记录
func UpdateAudioBook(id uint, updates map[string]interface{}) {
	var book models.AudioBook
	result := DB.Model(&book).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		log.Println("更新失败:", result.Error)
	} else {
		log.Println("更新成功:", updates)
	}
}

// 删除记录
func DeleteAudioBook(id uint) {
	var book models.AudioBook
	result := DB.Delete(&book, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}
