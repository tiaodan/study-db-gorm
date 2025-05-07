// db country 相关操作
package db

import (
	"log"
	"study-db-gorm/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func AddCountry(country *models.Country) error {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "NameId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": country.Name}),
	}).Create(country)
	if result.Error != nil {
		log.Println("创建失败:", result.Error)
		return result.Error
	} else {
		log.Println("创建成功:", country)
	}
	return nil
}

// 批量增
func BatchAddCountries(countries []*models.Country) {
	for i, country := range countries {
		err := AddCountry(country)
		if err == nil {
			log.Printf("批量创建第%d条成功, country: %v", i+1, &country)
		} else {
			log.Printf("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func DeleteCountry(id uint) {
	var country models.Country
	result := DB.Delete(&country, id)
	if result.Error != nil {
		log.Println("删除失败:", result.Error)
	} else {
		log.Println("删除成功:", id)
	}
}

// 批量删
func BatchDeleteCountries(ids []uint) {
	var countries []models.Country
	result := DB.Delete(&countries, ids)
	if result.Error != nil {
		log.Println("批量删除失败:", result.Error)
	} else {
		log.Println("批量删除成功:", ids)
	}
}

// 改
func UpdateCountry(nameId uint, updates map[string]interface{}) {
	var country models.Country
	// 解决0值不更新
	result := DB.Model(&country).Where("name_id = ?", nameId).Select("name").Updates(updates)
	if result.Error != nil {
		log.Println("修改失败:", result.Error)
	} else {
		log.Println("修改成功:", nameId)
	}
}

// 批量改
func BatchUpdateCountries(updates map[uint]map[string]interface{}) {
	for nameId, update := range updates {
		var country models.Country
		// 解决0值不更新
		result := DB.Model(&country).Where("name_id = ?", nameId).Select("name").Updates(update)
		if result.Error != nil {
			log.Printf("更新国家 %d 失败: %v\n", nameId, result.Error)
		} else {
			log.Printf("更新国家 %d 成功\n", nameId)
		}
	}
}

// 查
func QueryCountryById(id uint) *models.Country {
	var country models.Country
	result := DB.First(&country, id)
	if result.Error != nil {
		log.Println("查询失败:", result.Error)
		return nil
	}
	log.Println("查询成功:", country)
	return &country
}

// 批量查
func BatchQueryCountries(ids []uint) ([]*models.Country, error) {
	var countries []*models.Country
	result := DB.Find(&countries, ids)
	if result.Error != nil {
		log.Printf("批量查询失败: %v\n", result.Error)
		return countries, result.Error
	}
	log.Printf("批量查询成功, 查询到 %d 条记录\n", len(countries))
	return countries, nil
}
