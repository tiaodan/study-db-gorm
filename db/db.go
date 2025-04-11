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
	websiteDefaultJ88d := &models.Website{Name: "j88d", NameId: 1, URL: "www.j88d.com"}         // 请求url 时带上http://
	websiteDefaultXimalaya := &models.Website{Name: "喜马拉雅", NameId: 2, URL: "www.ximalaya.com"} // 请求url 时带上http://
	defaultWebsites := []*models.Website{websiteDefaultNoClass, websiteDefaultJ88d, websiteDefaultXimalaya}
	BatchAddWebsite(defaultWebsites)

	// 插入默认数据-class
	classDefaultNoClass := &models.Class{Name: "待分类", NameId: 0}
	classDefaultYouShengShu := &models.Class{Name: "有声书", NameId: 1}
	classDefaultGuangBoJu := &models.Class{Name: "广播剧", NameId: 2}
	classes := []*models.Class{classDefaultNoClass, classDefaultYouShengShu, classDefaultGuangBoJu}
	BatchAddClasses(classes)

	// 插入默认数据-country
	countryDefaultNoType := &models.Country{Name: "待分类", NameId: 0}
	countryDefaultChina := &models.Country{Name: "中国", NameId: 1}
	countryDefaultKoren := &models.Country{Name: "韩国", NameId: 2}
	countryDefaultAmerica := &models.Country{Name: "欧美", NameId: 3}
	countryDefaultJapan := &models.Country{Name: "日本", NameId: 4}
	countries := []*models.Class{countryDefaultNoType, countryDefaultChina, countryDefaultKoren, countryDefaultAmerica, countryDefaultJapan}
	BatchAddClasses(countries)

	// 插入默认数据-type
	/*
		- 一级分类： 多种相似类型的 组合。如玄幻奇幻、悬疑惊悚、武侠仙侠、都市言情、文学名著、科幻、穿越架空、军事历史、官场商战、游戏动漫、青春纯爱
		- 二级分类：
		玄幻奇幻：
		- 玄幻 √
		- 奇幻 √
		悬疑惊悚：
		- 悬疑  √
		- 惊悚 √
		武侠仙侠：
		- 武侠 √
		- 仙侠 √
		都市言情：
		- 都市 √
		- 言情 √
		文学名著： √
		科幻： √
		穿越架空：
		- 穿越  √
		- 架空     √
		军事历史：
		- 军事 √
		- 历史 √
		官场商战：
		- 官场 √
		- 商战 √
		游戏动漫：
		- 游戏 √
		- 动漫 √
		青春纯爱：
		- 青春 √
		- 纯爱 √

		- 三级分类
		悬疑惊悚：
		- 悬疑
		- 民间故事
		- 民间怪谈
		- 都市传说
		- 古风聊斋
		- 悬疑惊悚
		- 恐怖惊悚
		- 盗墓探险
		- 风水玄幻
		- 悬疑搞笑
		- 惊悚犯罪
		- 刑侦犯罪
		- 大案纪实
		- 都市刑侦
		- 犯罪现场
		- 侦探推理
		- 扫黑
		- 未解之谜
		- 世界未解之迷
		- 中国未解之谜
		武侠仙侠：
		- 武侠
		- 仙侠
		- 古典仙侠  是否要合并成一个？
		- 玄幻仙侠
		- 修真仙侠
		都市言情：
		- 都市
		- 言情
		- 现代言情
		- 古风言情
		游戏动漫：
		- 游戏
		- 电竞
		- 游戏
		- 动漫
		青春纯爱：
		- 青春
		- 纯爱
		- 古代纯爱
		- 现代纯爱
		- 同人纯爱
		穿越架空：
		- 穿越
		- 架空
		- 完全架空
		- 半架空
	*/

	// 一级分类
	// typeDefaultNoClass := &models.Type{NameId: 0, Name: "待分类", Level: 1}
	// typeDefaultYanqing := &models.Type{NameId: 1, Name: "言情", Level: 1}
	// typeDefaultXuanyi := &models.Type{NameId: 2, Name: "悬疑", Level: 1}
	// typeDefaultQihuan := &models.Type{NameId: 3, Name: "奇幻", Level: 1}
	// typeDefaultKehuan := &models.Type{NameId: 4, Name: "科幻", Level: 1}
	// typeDefaultXianxia := &models.Type{NameId: 5, Name: "仙侠", Level: 1}
	// typeDefaultWuxia := &models.Type{NameId: 6, Name: "武侠", Level: 1}
	// typeDefaultDushi := &models.Type{NameId: 7, Name: "都市", Level: 1}
	// typeDefaultJiakongChuanyue := &models.Type{NameId: 9, Name: "架空穿越", Level: 1}
	// typeDefaultWenXueMingzhu := &models.Type{NameId: 10, Name: "文学名著", Level: 1}
	// typeDefaultLishi := &models.Type{NameId: 11, Name: "历史", Level: 1}
	// typeDefaultJunshi := &models.Type{NameId: 12, Name: "军事", Level: 1}
	// typeDefaultYouxiDongman := &models.Type{NameId: 13, Name: "游戏动漫", Level: 1}
	// typeDefaultChunai := &models.Type{NameId: 14, Name: "纯爱", Level: 1}

	// // 二级分类
	// // 言情：111-119 参考NameId: 古代言情 现代言情
	// typeSubYanqingGudai := &models.Type{NameId: 111, Name: "古代言情", Level: 2, ParentId: 1}
	// typeSubYanqingXiandai := &models.Type{NameId: 112, Name: "现代言情", Level: 2, ParentId: 1}

	// // 悬疑：121-129 参考NameId: 推理探案 灵异恐怖
	// typeSubXuanyiTanan := &models.Type{NameId: 121, Name: "推理探案", Level: 2, ParentId: 2}
	// typeSubXuanyiKongbu := &models.Type{NameId: 122, Name: "灵异恐怖", Level: 2, ParentId: 2}

	// // 仙侠: 151-159 参考NameId: 古典仙侠 玄幻仙侠 修真仙侠
	// typeSubXianxiaGudian := &models.Type{NameId: 151, Name: "古典仙侠", Level: 2, ParentId: 5}
	// typeSubXianxiaXuanhuan := &models.Type{NameId: 152, Name: "玄幻仙侠", Level: 2, ParentId: 5}
	// typeSubXianxiaXiuzhen := &models.Type{NameId: 153, Name: "修真仙侠", Level: 2, ParentId: 5}

	// // 架空: 191-199 参考NameId: 完全架空 半架空
	// typeSubJiakongChuanyueWanquan := &models.Type{NameId: 191, Name: "完全架空", Level: 2, ParentId: 9}
	// typeSubJiakongChuanyueBanquan := &models.Type{NameId: 192, Name: "半架空", Level: 2, ParentId: 9}

	// // 纯爱: 241-249 参考NameId: 现代纯爱 古代纯爱 同人纯爱
	// typeSubChunaiXiandai := &models.Type{NameId: 241, Name: "现代纯爱", Level: 2, ParentId: 14}
	// typeSubChunaiGudai := &models.Type{NameId: 242, Name: "古代纯爱", Level: 2, ParentId: 14}
	// typeSubChunaiTongren := &models.Type{NameId: 243, Name: "同人纯爱", Level: 2, ParentId: 14}

	// // 游戏: 231-239 参考NameId: 电竞 动漫 游戏
	// typeSubYouxiDianjing := &models.Type{NameId: 231, Name: "电竞", Level: 2, ParentId: 13}
	// typeDefaultYouxiDongman := &models.Type{NameId: 232, Name: "动漫", Level: 13}
	// typeDefaultYouxiYouxi := &models.Type{NameId: 233, Name: "游戏", Level: 13}

	// defaultTypes := []*models.Type{
	// 	typeDefaultNoClass, typeDefaultYanqing, typeDefaultXuanyi,
	// 	typeDefaultQihuan, typeDefaultKehuan, typeDefaultXianxia,
	// 	typeDefaultWuxia, typeDefaultDushi, typeDefaultJiakongChuanyue,
	// 	typeDefaultWenXueMingzhu, typeDefaultLishi, typeDefaultJunshi,
	// 	typeDefaultYouxiDongman, typeDefaultChunai,
	// 	typeSubYanqingGudai, typeSubYanqingXiandai,
	// 	typeSubXuanyiTana, typeSubXuanyiKongbu,
	// 	typeSubChunaiXiandai, typeSubChunaiGudai, typeSubChunaiTongren,
	// }
	// BatchAddTypes(defaultTypes)

}
