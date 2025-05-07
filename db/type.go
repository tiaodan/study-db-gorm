// db type 相关操作
package db

import (
	"log"
	"study-db-gorm/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func AddType(typeData *models.Type) error {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "NameId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": typeData.Name, "level": typeData.Level, "parent": typeData.Parent}),
	}).Create(typeData)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
		return result.Error
	} else {
		log.Println("创建成功:", typeData)
	}
	return nil
}

// 批量增
func BatchAddTypes(types []*models.Type) {
	for i, typeData := range types {
		err := AddType(typeData)
		if err == nil {
			log.Printf("批量创建第%d条成功, type: %v", i+1, &typeData)
		} else {
			log.Printf("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func DeleteType(id uint) {
	var typeData models.Type
	result := DB.Delete(&typeData, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}

// 批量删
func BatchDeleteTypes(ids []uint) {
	var types []models.Type
	result := DB.Delete(&types, ids)
	if result.Error != nil {
		log.Println("批量删除失败:", result.Error)
	} else {
		log.Println("批量删除成功:", ids)
	}
}

// 改
func UpdateType(nameId uint, updates map[string]interface{}) {
	var typeData models.Type
	// 解决0值不更新问题
	result := DB.Model(&typeData).Where("name_id = ?", nameId).Select("name", "level", "parent").Updates(updates)
	if result.Error != nil {
		log.Println("修改失败:", result.Error)
	} else {
		log.Println("修改成功:", nameId)
	}
}

// 批量改
func BatchUpdateTypes(updates map[uint]map[string]interface{}) {
	for nameId, update := range updates {
		var typeData models.Type
		// 解决0值不更新问题
		result := DB.Model(&typeData).Where("name_id = ?", nameId).Select("name", "level", "parent").Updates(update)
		if result.Error != nil {
			log.Printf("更新类型 %d 失败: %v\n", nameId, result.Error)
		} else {
			log.Printf("更新类型 %d 成功\n", nameId)
		}
	}
}

// 查
func QueryTypeById(id uint) *models.Type {
	var typeData models.Type
	result := DB.First(&typeData, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", typeData)
	return &typeData
}

// 批量查
func BatchQueryTypes(ids []uint) ([]*models.Type, error) {
	var types []*models.Type
	result := DB.Find(&types, ids)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return types, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(types))
	return types, nil
}
