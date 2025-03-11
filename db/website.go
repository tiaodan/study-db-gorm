// db website 相关操作
package db

import (
	"log"
	"study-db-gorm/models"
)

// 增
func CreateWebsite(website *models.Website) {
	result := DB.Create(website)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
	} else {
		log.Println("创建成功:", website)
	}
}

// 删
func DeleteWebsite(id uint) {
	var website models.Website
	result := DB.Delete(&website, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}

// 改
func UpdateWebsite(id uint, updates map[string]interface{}) {
	var website models.Website
	DB.Model(&website).Where("id = ?", id).Updates(updates)
}

// 查
func QueryWebsiteById(id uint) *models.Website {
	var website models.Website
	result := DB.First(&website, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", website)
	return &website

}
