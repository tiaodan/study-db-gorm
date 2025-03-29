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
	db.DB.AutoMigrate(&models.Website{}, &models.Type{}, &models.Class{}) // 有几个表, 写几个参数

	// 插入默认数据
	db.InsertDefaultData()
}

func main() {

	// 插入website默认数据
	// websiteDefaultNoClass := &models.Website{Name: "待分类", NameId: 10, URL: "11122"} // 未分类
	// websiteDefaultJ88d := &models.Website{Name: "j88d", NameId: 1, URL: "http://www.j88d.com"}
	// websites := []*models.Website{websiteDefaultNoClass, websiteDefaultJ88d}
	// db.BatchAddWebsite(websites)

	// db.AddWebsite(websiteDefaultNoClass)

	// db.DeleteWebsite(1)
	// websiteIds := []uint{7, 8, 12, 15, 17}
	// db.BatchDeleteWebsites(websiteIds)

	// 改
	// db.UpdateWebsite(2, map[string]interface{}{"Name": "待分类111", "NameId": 2, "URL": "未知111"})
	// 批量改
	// updates := map[uint]map[string]interface{}{
	// 	0: {"Name": "待分类111", "NameId": 0, "URL": "未知111"},
	// 	1: {"Name": "待分类222", "NameId": 1, "URL": "未知222"},
	// }

	// db.BatchUpdateWebsites(updates)

	// 查
	// db.QueryWebsiteById(3)
	// 批量查
	// db.BatchQueryWebsites([]uint{1, 2, 3, 1, 5})

	// ----------------------测试数据 - type表
	// 增
	// typeDefaultNoClass := &models.Type{Name: "待分类11", NameId: 0, Level: 111, Parent: 111} // 未分类
	// db.AddType(typeDefaultNoClass)

	// 批量增
	// typeDefaultNoClass := &models.Type{Name: "待分类11", NameId: 0, Level: 111, Parent: 111} // 未分类
	// typeDefaultXuanhuan := &models.Type{Name: "玄幻11", NameId: 1, Level: 111, Parent: 111} // 未分类
	// types := []*models.Type{typeDefaultNoClass, typeDefaultXuanhuan}
	// db.BatchAddTypes(types)

	// 删
	// db.DeleteType(10)

	// 批量删
	// typeIds := []uint{8, 39}
	// db.BatchDeleteTypes(typeIds)

	// 改
	// db.UpdateType(3, map[string]interface{}{"Name": "待分类122", "Level": 3})

	// 批量改
	// updates := map[uint]map[string]interface{}{
	// 	0: {"Name": "待分类111"},
	// 	1: {"Name": "待分类222"},
	// }
	// db.BatchUpdateTypes(updates)

	// 查
	// db.QueryTypeById(40)

	// 批量查
	// db.BatchQueryTypes([]uint{8, 10, 40})

	// ----------------------测试数据 - class表
	// 增
	// classDefaultNoClass := &models.Class{Name: "待分类", NameId: 0} // 未分类
	// db.AddClass(classDefaultNoClass)

	// 批量增
	// classDefaultNoClass := &models.Class{Name: "待分类", NameId: 0}  // 未分类
	// classDefaultXuanhuan := &models.Class{Name: "有声书", NameId: 1} // 未分类
	// classes := []*models.Class{classDefaultNoClass, classDefaultXuanhuan}
	// db.BatchAddClasses(classes)

	// 删
	// db.DeleteClass(1)

	// 批量删
	// classIds := []uint{3, 6}
	// db.BatchDeleteClasses(classIds)

	// 改
	// db.UpdateClass(0, map[string]interface{}{"Name": "待分类122", "NameId": 0})

	// 批量改
	// updates := map[uint]map[string]interface{}{
	// 	0: {"Name": "待分类111"},
	// 	1: {"Name": "待分类222"},
	// }
	// db.BatchUpdateClasses(updates)

	// 查
	// db.QueryClassById(8)

	// 批量查
	// db.BatchQueryClasses([]uint{8, 9, 11})
}
