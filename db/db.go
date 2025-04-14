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
	classDefaultPingshu := &models.Class{Name: "评书", NameId: 3}
	classes := []*models.Class{classDefaultNoClass, classDefaultYouShengShu, classDefaultGuangBoJu, classDefaultPingshu}
	BatchAddClasses(classes)

	// 插入默认数据-country
	countryDefaultNoType := &models.Country{Name: "待分类", NameId: 0}
	countryDefaultChina := &models.Country{Name: "中国", NameId: 1}
	countryDefaultKoren := &models.Country{Name: "韩国", NameId: 2}
	countryDefaultAmerica := &models.Country{Name: "欧美", NameId: 3}
	countryDefaultJapan := &models.Country{Name: "日本", NameId: 4}
	countries := []*models.Country{countryDefaultNoType, countryDefaultChina, countryDefaultKoren, countryDefaultAmerica, countryDefaultJapan}
	BatchAddCountries(countries)

	// 插入默认数据-type
	/*
		- 一级分类： 多种相似类型的 组合。如待分类、玄幻奇幻、悬疑惊悚、武侠仙侠、都市言情、文学名著、科幻、穿越架空、军事历史、官场商战、游戏动漫、青春纯爱
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
	typeDefaultNoTypeLevel1 := &models.Type{NameId: 0, Name: "待分类", Level: 1}
	typeDefaultXuanhuanQihuan := &models.Type{NameId: 1, Name: "玄幻奇幻", Level: 1}
	typeDefaultXuanyiJingsong := &models.Type{NameId: 2, Name: "悬疑惊悚", Level: 1}
	typeDefaultWuxiaXianxia := &models.Type{NameId: 3, Name: "武侠仙侠", Level: 1}
	typeDefaultDushiYanqing := &models.Type{NameId: 4, Name: "都市言情", Level: 1}
	typeDefaultWenXueMingzhu := &models.Type{NameId: 5, Name: "文学名著", Level: 1}
	typeDefaultKehuan := &models.Type{NameId: 6, Name: "科幻", Level: 1}
	typeDefaultChuanyueiakong := &models.Type{NameId: 7, Name: "穿越架空", Level: 1}
	typeDefaultJunshiLishi := &models.Type{NameId: 9, Name: "军事历史", Level: 1}
	typeDefaultGuanchangShangzhan := &models.Type{NameId: 10, Name: "官场商战", Level: 1}
	typeDefaultYouxiDongman := &models.Type{NameId: 11, Name: "游戏动漫", Level: 1}
	typeDefaultQingchunChunai := &models.Type{NameId: 12, Name: "青春纯爱", Level: 1}

	// 二级分类
	// 未分类 leve2
	typeSubNoTypeLevel2 := &models.Type{NameId: 13, Name: "待分类", Level: 2}

	// 玄幻奇幻
	typeSubXuanhuan := &models.Type{NameId: 14, Name: "玄幻", Level: 2, Parent: 1}
	typeSubQihuan := &models.Type{NameId: 15, Name: "奇幻", Level: 2, Parent: 1}

	// 悬疑惊悚
	typeSubXuanyi := &models.Type{NameId: 16, Name: "悬疑", Level: 2, Parent: 2}
	typeSubJingsong := &models.Type{NameId: 17, Name: "惊悚", Level: 2, Parent: 2}

	// 武侠仙侠
	typeSubWuxia := &models.Type{NameId: 18, Name: "武侠", Level: 2, Parent: 3}
	typeSubXianxia := &models.Type{NameId: 19, Name: "仙侠", Level: 2, Parent: 3}

	// 都市言情
	typeSubDushi := &models.Type{NameId: 20, Name: "都市", Level: 2, Parent: 4}
	typeSubYanqing := &models.Type{NameId: 21, Name: "言情", Level: 2, Parent: 4}

	// 穿越架空
	typeSubChuanyue := &models.Type{NameId: 22, Name: "穿越", Level: 2, Parent: 7}
	typeSubJiakong := &models.Type{NameId: 23, Name: "架空", Level: 2, Parent: 7}

	// 军事历史
	typeSubJunshi := &models.Type{NameId: 24, Name: "军事", Level: 2, Parent: 9}
	typeSubLishi := &models.Type{NameId: 25, Name: "历史", Level: 2, Parent: 9}

	// 官场商战
	typeSubGuanchang := &models.Type{NameId: 26, Name: "官场", Level: 2, Parent: 10}
	typeSubShangzhan := &models.Type{NameId: 27, Name: "商战", Level: 2, Parent: 10}

	// 游戏动漫
	typeSubYouxi := &models.Type{NameId: 28, Name: "游戏", Level: 2, Parent: 11}
	typeSubDongman := &models.Type{NameId: 29, Name: "动漫", Level: 2, Parent: 11}

	// 青春纯爱
	typeSubQingchun := &models.Type{NameId: 30, Name: "青春", Level: 2, Parent: 12}
	typeSubChunai := &models.Type{NameId: 31, Name: "纯爱", Level: 2, Parent: 12}

	// 三级分类
	/*
		- 三级分类
			悬疑惊悚：
			- 悬疑
			- 民间故事：
				- 民间怪谈
				- 都市传说
				- 古风聊斋
			- 悬疑惊悚：
				- 恐怖惊悚
				- 盗墓探险
				- 风水玄幻
				- 悬疑搞笑
				- 惊悚犯罪
			- 刑侦犯罪:
				- 大案纪实
				- 都市刑侦
				- 犯罪现场
				- 侦探推理
				- 扫黑
			- 未解之谜:
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

	// 仙侠-3级
	typeSubXianxiaGudian := &models.Type{NameId: 32, Name: "古典仙侠", Level: 3, Parent: 19}
	typeSubXianxiaXuanhuan := &models.Type{NameId: 33, Name: "玄幻仙侠", Level: 3, Parent: 19}
	typeSubXianxiaXiuzhen := &models.Type{NameId: 34, Name: "修真仙侠", Level: 3, Parent: 19}

	// 言情-3级
	typeSubYanqingXiandai := &models.Type{NameId: 35, Name: "现代言情", Level: 3, Parent: 21}
	typeSubYanqingGufeng := &models.Type{NameId: 36, Name: "古风言情", Level: 3, Parent: 21}

	// 游戏-3级
	typeSubYouxiLevel3 := &models.Type{NameId: 37, Name: "游戏", Level: 3, Parent: 28}
	typeSubYouxiDianjing := &models.Type{NameId: 38, Name: "电竞", Level: 3, Parent: 28}

	// 纯爱-3级
	typeSubChunaiGudai := &models.Type{NameId: 39, Name: "古代纯爱", Level: 3, Parent: 31}
	typeSubChunaiXiandai := &models.Type{NameId: 40, Name: "现代纯爱", Level: 3, Parent: 31}
	typeSubChunaiTongren := &models.Type{NameId: 41, Name: "同人纯爱", Level: 3, Parent: 31}

	// 架空-3级
	typeSubJiaKongWanquan := &models.Type{NameId: 42, Name: "完全架空", Level: 3, Parent: 22}
	typeSubJiakongBan := &models.Type{NameId: 43, Name: "半架空", Level: 3, Parent: 22}

	// 悬疑
	typeSubMinjianGushi := &models.Type{NameId: 44, Name: "民间故事", Level: 3, Parent: 16}
	typeSubXuanyiJingsong := &models.Type{NameId: 45, Name: "悬疑惊悚", Level: 3, Parent: 16}
	typeSubXingzhenFanzui := &models.Type{NameId: 46, Name: "刑侦犯罪", Level: 3, Parent: 16}
	typeSubWeijieZhimi := &models.Type{NameId: 47, Name: "未解之谜", Level: 3, Parent: 16}

	// 四级分类-----------
	// - 民间故事
	typeSubMinjianGuaitan := &models.Type{NameId: 48, Name: "民间怪谈", Level: 4, Parent: 44}
	typeSubDushiChuanshuo := &models.Type{NameId: 49, Name: "都市传说", Level: 4, Parent: 44}
	typeSubGufengLiaozhai := &models.Type{NameId: 50, Name: "古风聊斋", Level: 4, Parent: 44}
	// - 悬疑惊悚
	typeSubKongbuJingsong := &models.Type{NameId: 51, Name: "恐怖惊悚", Level: 4, Parent: 45}
	typeSubDaomuTanxian := &models.Type{NameId: 52, Name: "盗墓探险", Level: 4, Parent: 45}
	typeSubFengshuiXuanhuan := &models.Type{NameId: 53, Name: "风水玄幻", Level: 4, Parent: 45}
	typeSubXuanyiGaoxiao := &models.Type{NameId: 54, Name: "悬疑搞笑", Level: 4, Parent: 45}
	typeSubJiongsongFanzui := &models.Type{NameId: 55, Name: "惊悚犯罪", Level: 4, Parent: 45}
	// - 刑侦犯罪
	typeSubDaanJishi := &models.Type{NameId: 56, Name: "大案纪实", Level: 4, Parent: 46}
	typeSubDushiXingzhen := &models.Type{NameId: 57, Name: "都市刑侦", Level: 4, Parent: 46}
	typeSubFanzuiXianchang := &models.Type{NameId: 58, Name: "犯罪现场", Level: 4, Parent: 46}
	typeSubXingzhenTuili := &models.Type{NameId: 59, Name: "刑侦推理", Level: 4, Parent: 46}
	typeSubSaohei := &models.Type{NameId: 60, Name: "扫黑", Level: 4, Parent: 46}
	// - 未解之谜
	typeSubWeijieZhimiShijie := &models.Type{NameId: 61, Name: "世界未解之谜", Level: 4, Parent: 47}
	typeSubWeijieZhimiZhongguo := &models.Type{NameId: 62, Name: "中国未解之谜", Level: 4, Parent: 47}

	defaultTypes := []*models.Type{
		// 一级分类
		typeDefaultNoTypeLevel1, typeDefaultXuanhuanQihuan, typeDefaultXuanyiJingsong,
		typeDefaultWuxiaXianxia, typeDefaultDushiYanqing, typeDefaultWenXueMingzhu,
		typeDefaultKehuan, typeDefaultChuanyueiakong, typeDefaultJunshiLishi,
		typeDefaultGuanchangShangzhan, typeDefaultYouxiDongman, typeDefaultQingchunChunai,
		// 二级分类
		typeSubNoTypeLevel2,
		typeSubXuanhuan, typeSubQihuan,
		typeSubXuanyi, typeSubJingsong,
		typeSubWuxia, typeSubXianxia,
		typeSubDushi, typeSubYanqing,
		typeSubChuanyue, typeSubJiakong,
		typeSubJunshi, typeSubLishi,
		typeSubGuanchang, typeSubShangzhan,
		typeSubYouxi, typeSubDongman,
		typeSubQingchun, typeSubChunai,
		// 三级分类
		typeSubXianxiaGudian, typeSubXianxiaXuanhuan, typeSubXianxiaXiuzhen,
		typeSubYanqingXiandai, typeSubYanqingGufeng,
		typeSubYouxiLevel3, typeSubYouxiDianjing,
		typeSubChunaiGudai, typeSubChunaiXiandai, typeSubChunaiTongren,
		typeSubJiaKongWanquan, typeSubJiakongBan,
		typeSubMinjianGushi, typeSubXuanyiJingsong, typeSubXingzhenFanzui, typeSubWeijieZhimi,
		// 四级分类
		typeSubMinjianGuaitan, typeSubDushiChuanshuo, typeSubGufengLiaozhai,
		typeSubKongbuJingsong, typeSubDaomuTanxian, typeSubFengshuiXuanhuan, typeSubXuanyiGaoxiao, typeSubJiongsongFanzui,
		typeSubDaanJishi, typeSubDushiXingzhen, typeSubFanzuiXianchang, typeSubXingzhenTuili, typeSubSaohei,
		typeSubWeijieZhimiShijie, typeSubWeijieZhimiZhongguo,
	}

	BatchAddTypes(defaultTypes)

}
