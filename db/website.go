// db website 相关操作
package db

import (
	"log"
	"study-db-gorm/models"
)

// 增
func AddWebsite(website *models.Website) {
	result := DB.Create(website)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
	} else {
		log.Println("创建成功:", website)
	}
}

// 批量增
func BatchAddWebsite(websites []*models.Website) {
	result := DB.Create(websites)
	if result.Error != nil {
		log.Println("批量创建失败:", result.Error)
	} else {
		log.Println("批量创建成功:", websites)
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

// 批量删
func BatchDeleteWebsites(ids []uint) {
	var websites []models.Website
	result := DB.Delete(&websites, ids)
	if result.Error != nil {
		log.Println("批量删除失败:", result.Error)
	} else {
		log.Println("批量删除成功:", ids)
	}
}

// 改
func UpdateWebsite(id uint, updates map[string]interface{}) {
	var website models.Website
	DB.Model(&website).Where("id = ?", id).Updates(updates)
}

// 批量改
func BatchUpdateWebsites(updates map[uint]map[string]interface{}) {
	for id, update := range updates {
		var website models.Website
		result := DB.Model(&website).Where("id = ?", id).Updates(update)
		if result.Error != nil {
			log.Printf("更新网站 %d 失败: %v\n", id, result.Error)
		} else {
			log.Printf("更新网站 %d 成功\n", id)
		}
	}
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
