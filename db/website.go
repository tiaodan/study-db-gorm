// db website 相关操作
package db

import (
	"log"
	"study-db-gorm/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func AddWebsite(website *models.Website) error {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "NameId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": website.Name, "url": website.URL}),
	}).Create(website)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
		return result.Error
	} else {
		log.Println("创建成功:", website)
	}
	return nil
}

// 批量增
func BatchAddWebsite(websites []*models.Website) {
	for i, website := range websites {
		err := AddWebsite(website)
		if err == nil {
			log.Printf("批量创建第%d条成功, website: %v", i+1, &website)
		} else {
			log.Printf("批量创建第%d条失败, err: %v", i+1, err)
		}
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
func UpdateWebsite(nameId uint, updates map[string]interface{}) {
	var website models.Website
	result := DB.Model(&website).Where("name_id = ?", nameId).Updates(updates)
	if result.Error != nil {
		log.Println("修改失败:", result.Error)
	} else {
		log.Println("修改成功:", nameId)
	}
}

// 批量改
func BatchUpdateWebsites(updates map[uint]map[string]interface{}) {
	for nameId, update := range updates {
		var website models.Website
		result := DB.Model(&website).Where("name_id = ?", nameId).Updates(update)
		if result.Error != nil {
			log.Printf("更新网站 %d 失败: %v\n", nameId, result.Error)
		} else {
			log.Printf("更新网站 %d 成功\n", nameId)
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

// 批量查
func BatchQueryWebsites(ids []uint) ([]*models.Website, error) {
	var websites []*models.Website
	result := DB.Find(&websites, ids)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return websites, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(websites))
	return websites, nil
}
