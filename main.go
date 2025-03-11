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

	// 插入website默认数据
	// websiteDefaultNoClass := &models.Website{Name: "待分类", NameId: 0, URL: "未知"} // 未分类
	// websiteDefaultJ88d := &models.Website{Name: "j88d", NameId: 1, URL: "http://www.j88d.com"}
	// websites := []*models.Website{websiteDefaultNoClass, websiteDefaultJ88d}
	// db.BatchAddWebsite(websites)

	// db.DeleteWebsite(1)
	// websiteIds := []uint{5, 6}
	// db.BatchDeleteWebsites(websiteIds)

	// 改
	// db.UpdateWebsite(2, map[string]interface{}{"Name": "待分类111", "NameId": 5, "URL": "未知111"})
	// 批量改
	updates := map[uint]map[string]interface{}{
		7: {"Name": "待分类111", "NameId": 7, "URL": "未知111"},
		8: {"Name": "待分类222", "NameId": 8, "URL": "未知222"},
	}

	db.BatchUpdateWebsites(updates)
	// db.QueryWebsiteById(3)

}
