package main

import (
	"log"
	"study-db-gorm/db"
	"study-db-gorm/models"
)

// 初始化
/*
思路:
1 设置日志格式：日期时间 + 短文件名 + 行号
2 初始化数据库连接
3 自动迁移表结构
*/
func init() {
	// 设置日志格式：日期时间 + 短文件名 + 行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化数据库连接
	db.InitDB()

	// 自动迁移表结构
	db.DB.AutoMigrate(&models.AudioBook{}, &models.Website{}) // 修改: 添加 &models.Website{}
}

func main() {

	// // 创建记录
	// book := &models.AudioBook{Title: "Go语言入门", Author: "张三"}
	// db.CreateAudioBook(book)

	// // 查询记录
	// retrievedBook := db.GetAudioBookByID(1)
	// if retrievedBook != nil {
	// 	fmt.Println("查询到的书籍:", retrievedBook)
	// }

	// // 更新记录
	// db.UpdateAudioBook(1, map[string]interface{}{"Title": "Go语言进阶"})

	// // 删除记录
	// db.DeleteAudioBook(1)

}
