// db class 相关操作
package db

import (
	"log"
	"study-db-gorm/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func AddClass(classData *models.Class) error {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "NameId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": classData.Name}),
	}).Create(classData)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
		return result.Error
	} else {
		log.Println("创建成功:", classData)
	}
	return nil
}

// 批量增
func BatchAddClasses(classes []*models.Class) {
	for i, classData := range classes {
		err := AddClass(classData)
		if err == nil {
			log.Printf("批量创建第%d条成功, class: %v", i+1, &classData)
		} else {
			log.Printf("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func DeleteClass(id uint) {
	var classData models.Class
	result := DB.Delete(&classData, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}

// 批量删
func BatchDeleteClasses(ids []uint) {
	var classes []models.Class
	result := DB.Delete(&classes, ids)
	if result.Error != nil {
		log.Println("批量删除失败:", result.Error)
	} else {
		log.Println("批量删除成功:", ids)
	}
}

// 改
func UpdateClass(nameId uint, updates map[string]interface{}) {
	var classData models.Class
	result := DB.Model(&classData).Where("name_id = ?", nameId).Updates(updates)
	if result.Error != nil {
		log.Println("修改失败:", result.Error)
	} else {
		log.Println("修改成功:", nameId)
	}
}

// 批量改
func BatchUpdateClasses(updates map[uint]map[string]interface{}) {
	for nameId, update := range updates {
		var classData models.Class
		result := DB.Model(&classData).Where("name_id = ?", nameId).Updates(update)
		if result.Error != nil {
			log.Printf("更新类型 %d 失败: %v\n", nameId, result.Error)
		} else {
			log.Printf("更新类型 %d 成功\n", nameId)
		}
	}
}

// 查
func QueryClassById(id uint) *models.Class {
	var classData models.Class
	result := DB.First(&classData, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", classData)
	return &classData
}

// 批量查
func BatchQueryClasses(ids []uint) ([]*models.Class, error) {
	var classes []*models.Class
	result := DB.Find(&classes, ids)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return classes, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(classes))
	return classes, nil
}
