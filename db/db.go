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

// 插入默认数据
/*
思路:
1. 插入默认数据-website
2. 插入默认数据-type
*/
func InsertDefaultData() {
	// 插入默认数据-website
	websiteDefaultNoClass := &models.Website{Name: "待分类", NameId: 0, URL: "未知"}
	websiteDefaultJ88d := &models.Website{Name: "j88d", NameId: 1, URL: "http://www.j88d.com"}
	defaultWebsites := []*models.Website{websiteDefaultNoClass, websiteDefaultJ88d}
	BatchAddWebsite(defaultWebsites)

	// 插入默认数据-class
	classDefaultNoClass := &models.Class{Name: "待分类", NameId: 0}
	classDefaultYouShengShu := &models.Class{Name: "有声书", NameId: 1}
	classDefaultGuangBoJu := &models.Class{Name: "广播剧", NameId: 2}
	classes := []*models.Class{classDefaultNoClass, classDefaultYouShengShu, classDefaultGuangBoJu}
	BatchAddClasses(classes)

	// 插入默认数据-type
	/*
		悬疑分：
		推理探案 灵异恐怖

		言情：
		古代言情 现代言情

		纯爱:
		现代纯爱 古代纯爱 同人纯爱

		游戏:
		电竞 动漫 游戏

		仙侠:
		古典仙侠 玄幻仙侠 修真仙侠

		架空:
		完全架空 半架空

		科幻:
		武侠:
		奇幻:
		都市:
		文学名著:

	*/
	typeDefaultNoClass := &models.Type{NameId: 0, Name: "待分类", Level: 1}
	typeDefaultYanqing := &models.Type{NameId: 1, Name: "言情", Level: 1}
	typeDefaultXuanyi := &models.Type{NameId: 2, Name: "悬疑", Level: 1}
	typeDefaultQihuan := &models.Type{NameId: 3, Name: "奇幻", Level: 1}
	typeDefaultKehuan := &models.Type{NameId: 4, Name: "科幻", Level: 1}
	typeDefaultXianxia := &models.Type{NameId: 5, Name: "仙侠", Level: 1}
	typeDefaultWuxia := &models.Type{NameId: 6, Name: "武侠", Level: 1}
	typeDefaultDushi := &models.Type{NameId: 7, Name: "都市", Level: 1}
	typeDefaultJiakongChuanyue := &models.Type{NameId: 9, Name: "架空穿越", Level: 1}
	typeDefaultWenXueMingzhu := &models.Type{NameId: 10, Name: "文学名著", Level: 1}
	typeDefaultLishi := &models.Type{NameId: 11, Name: "历史", Level: 1}
	typeDefaultJunshi := &models.Type{NameId: 12, Name: "军事", Level: 1}
	typeDefaultYouxiDongman := &models.Type{NameId: 13, Name: "游戏动漫", Level: 1}
	typeDefaultChunai := &models.Type{NameId: 14, Name: "纯爱", Level: 1}

	defaultTypes := []*models.Type{
		typeDefaultNoClass, typeDefaultYanqing, typeDefaultXuanyi,
		typeDefaultQihuan, typeDefaultKehuan, typeDefaultXianxia,
		typeDefaultWuxia, typeDefaultDushi, typeDefaultJiakongChuanyue,
		typeDefaultWenXueMingzhu, typeDefaultLishi, typeDefaultJunshi,
		typeDefaultYouxiDongman, typeDefaultChunai,
	}

	BatchAddTypes(defaultTypes)

}
