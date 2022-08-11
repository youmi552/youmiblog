package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false, //是否跳过事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_", //创建表添加前缀,User 会变为 t_users
			SingularTable: true,    //将表以单数形式存在, t_users 会变为t_user
		},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键(代码里自动体现外键关系)
	})
	if err != nil {
		panic(err)
	}
	db = con
	//M := db.Migrator()
	//if !M.HasTable(&models.Category{}) {
	//	M.CreateTable(&models.Category{})
	//	db.Create(&models.Category{Name: "Go"})
	//	db.Create(&models.Category{Name: "Java"})
	//}
	//db.AutoMigrate(&models.Category{}, &models.PostMore{})
	//db.Create(&models.Category{Name: "Go", PostMore: []models.PostMore{
	//	models.PostMore{
	//		Title:    "go博客实战",
	//		Content:  "牛嘛呀！",
	//		UserName: "悠米",
	//	},
	//	models.PostMore{
	//		Title:    "go博客实战2",
	//		Content:  "<blockquote>\n<p>现今互联网已经深入人们的生活，各种知识爆炸式袭来，随着行业的成熟以及学历人群的大幅度提升，随之而来的要求也越发提高，如果想要更加深入的学习提升，在激烈的竞争环境中脱颖而出，知识付费成为一个新的选择，知识是有价的，核心的高质量的课程和服务永远不可能免费。</p>\n</blockquote>\n<p>针对现今行业的现状，程序员面临的三个大问题：</p>\n<ul>\n<li><p><strong>学习</strong></p>\n<p>学习上面临的问题是无人指导，解决问题一靠百度，二靠问人，三靠自己琢磨，虽然最终可能会解决，但耗费了大量的时间，同时问题虽然解决了，但有时候根本不知道为什么？那么解决问题，还有意义吗？我将这种情况称为<code>无效学习</code>。</p>\n</li><li><p><strong>进阶</strong></p>\n<p>因为现在的岗位要求越来越高，针对的技能亮点，也要求更加深入，不管是实习，校招，社招遇到的问题都一样，CRUD的项目或者是无任何业务思考的项目能拿的出手吗？显然是不能。</p>\n<p>不管是初入职场还是工作多年的码农，都有一颗进步的心，可是职场毕竟是职场，不一定有核心的业务教给你做，然后就形成了一种恶性循环，没有核心知识和认知无法找到让自身成长的工作，而接触不到核心的业务又无法让自身掌握核心竞争力。</p>\n<p>上述的这些情况，我称为是<code>迷途的羔羊</code>。</p>\n</li><li><p><strong>面试</strong></p>\n<p>面试造火箭，工作拧螺丝，这是大家都清楚知道的，不管你技术多么厉害，面试过不去，一切都扯淡。</p>\n<p>面试靠什么？</p>\n<p>一是表达，流畅的表达，清晰的思路，给面试官留下的印象绝对深刻。</p>\n<p>二是简历，简历是获得面试机会的唯一入口，如何写一个好的简历，关乎生死。</p>\n<p>三是八股文，八股文都不会，自然不会更深入的和你谈，面试官的时间也很金贵。</p>\n<p>四是真实经验，没有实际做过，你吹牛逼的时候，面试官是否总露出一丝淡淡的笑意看着你。</p>\n<p>机会总是很有限，准备不足，面试效果自然很差，或者有人心态直接就崩了。</p>\n<p>这种情况我称为<code>战前准备不足</code></p>\n</li></ul>\n<blockquote>\n<p>针对以上情况，提供了一些核心服务</p>\n</blockquote>\n<ul>\n<li><strong>会员服务(一年期)</strong><ul>\n<li>任意问题的答疑指导服务，一对一，24小时在线的导师，解答满意为止</li><li>任意问题举例，比如日常学习疑惑，bug解决，工作问题，简历修改，面试指导，offer选择等等</li><li>详情介绍请看<a href=\"/pay/vip.md\">会员服务详细介绍</a>了解</li></ul>\n</li><li><strong>进阶付费课程</strong><ul>\n<li>码神学堂（真实项目实战教程，无惧面试，非DEMO项目）详情介绍请看<a href=\"/pay/xt.md\">码神学堂项目实战教程详细介绍</a>了解</li><li>码神RPC（从零开发一个RPC框架） 详情介绍请看<a href=\"/pay/rpc.md\">码神RPC详细介绍</a>了解</li><li>码神Spring（从零开发一个Spring框架）详情介绍请看<a href=\"/pay/spring.md\">码神Spring详细介绍</a>了解</li><li>架构师课程-全面认知技术世界（制作中）详情介绍请看<a href=\"/pay/svip.md\">架构师课程详细介绍</a>了解<ul>\n<li>第一篇章-全面认识计算机</li><li>第二篇章-全面认识语言</li><li>第三篇章-全面认识架构</li></ul>\n</li><li>Go游戏实战教程（制作中）详情介绍请看<a href=\"/pay/game.md\">go游戏实战教程详细介绍</a>了解</li><li>码神MQ（从零开发MQ中间件 制作中）</li><li>能造轮子的尽量都造一遍…</li><li>如果想多个套餐一起买，可以看<a href=\"/pay/tc.md\">套餐组合优惠详细介绍</a></li></ul>\n</li><li><strong>终身会员(即将涨价)</strong><ul>\n<li>买断制，所有已出和未出的付费课程和项目</li><li>终身会员服务，终身答疑指导</li><li>一对一模拟面试</li><li>等等</li></ul>\n</li><li><strong>go开发Offer收割社群</strong> <ul>\n<li>Go全套资料，包含学习，面试真题等等</li><li>一对一模拟面试，多次</li><li>模拟面试通过可获内推，跟踪整个流程</li><li>终身服务</li><li>详情介绍请看<a href=\"/pay/go.md\">go开发offer收割社群详细介绍</a>了解</li></ul>\n</li></ul>\n<blockquote>\n<p>如有意向，可联系我了解详情</p>\n</blockquote>\n<p><img src=\"./img/image-20220215214357342.png\" alt=\"image-20220215214357342\"></p>\n",
	//		UserName: "悠米",
	//	},
	//	models.PostMore{
	//		Title:    "go博客实战3",
	//		Content:  "<blockquote>\n<p>现今互联网已经深入人们的生活，各种知识爆炸式袭来，随着行业的成熟以及学历人群的大幅度提升，随之而来的要求也越发提高，如果想要更加深入的学习提升，在激烈的竞争环境中脱颖而出，知识付费成为一个新的选择，知识是有价的，核心的高质量的课程和服务永远不可能免费。</p>\n</blockquote>\n<p>针对现今行业的现状，程序员面临的三个大问题：</p>\n<ul>\n<li><p><strong>学习</strong></p>\n<p>学习上面临的问题是无人指导，解决问题一靠百度，二靠问人，三靠自己琢磨，虽然最终可能会解决，但耗费了大量的时间，同时问题虽然解决了，但有时候根本不知道为什么？那么解决问题，还有意义吗？我将这种情况称为<code>无效学习</code>。</p>\n</li><li><p><strong>进阶</strong></p>\n<p>因为现在的岗位要求越来越高，针对的技能亮点，也要求更加深入，不管是实习，校招，社招遇到的问题都一样，CRUD的项目或者是无任何业务思考的项目能拿的出手吗？显然是不能。</p>\n<p>不管是初入职场还是工作多年的码农，都有一颗进步的心，可是职场毕竟是职场，不一定有核心的业务教给你做，然后就形成了一种恶性循环，没有核心知识和认知无法找到让自身成长的工作，而接触不到核心的业务又无法让自身掌握核心竞争力。</p>\n<p>上述的这些情况，我称为是<code>迷途的羔羊</code>。</p>\n</li><li><p><strong>面试</strong></p>\n<p>面试造火箭，工作拧螺丝，这是大家都清楚知道的，不管你技术多么厉害，面试过不去，一切都扯淡。</p>\n<p>面试靠什么？</p>\n<p>一是表达，流畅的表达，清晰的思路，给面试官留下的印象绝对深刻。</p>\n<p>二是简历，简历是获得面试机会的唯一入口，如何写一个好的简历，关乎生死。</p>\n<p>三是八股文，八股文都不会，自然不会更深入的和你谈，面试官的时间也很金贵。</p>\n<p>四是真实经验，没有实际做过，你吹牛逼的时候，面试官是否总露出一丝淡淡的笑意看着你。</p>\n<p>机会总是很有限，准备不足，面试效果自然很差，或者有人心态直接就崩了。</p>\n<p>这种情况我称为<code>战前准备不足</code></p>\n</li></ul>\n<blockquote>\n<p>针对以上情况，提供了一些核心服务</p>\n</blockquote>\n<ul>\n<li><strong>会员服务(一年期)</strong><ul>\n<li>任意问题的答疑指导服务，一对一，24小时在线的导师，解答满意为止</li><li>任意问题举例，比如日常学习疑惑，bug解决，工作问题，简历修改，面试指导，offer选择等等</li><li>详情介绍请看<a href=\"/pay/vip.md\">会员服务详细介绍</a>了解</li></ul>\n</li><li><strong>进阶付费课程</strong><ul>\n<li>码神学堂（真实项目实战教程，无惧面试，非DEMO项目）详情介绍请看<a href=\"/pay/xt.md\">码神学堂项目实战教程详细介绍</a>了解</li><li>码神RPC（从零开发一个RPC框架） 详情介绍请看<a href=\"/pay/rpc.md\">码神RPC详细介绍</a>了解</li><li>码神Spring（从零开发一个Spring框架）详情介绍请看<a href=\"/pay/spring.md\">码神Spring详细介绍</a>了解</li><li>架构师课程-全面认知技术世界（制作中）详情介绍请看<a href=\"/pay/svip.md\">架构师课程详细介绍</a>了解<ul>\n<li>第一篇章-全面认识计算机</li><li>第二篇章-全面认识语言</li><li>第三篇章-全面认识架构</li></ul>\n</li><li>Go游戏实战教程（制作中）详情介绍请看<a href=\"/pay/game.md\">go游戏实战教程详细介绍</a>了解</li><li>码神MQ（从零开发MQ中间件 制作中）</li><li>能造轮子的尽量都造一遍…</li><li>如果想多个套餐一起买，可以看<a href=\"/pay/tc.md\">套餐组合优惠详细介绍</a></li></ul>\n</li><li><strong>终身会员(即将涨价)</strong><ul>\n<li>买断制，所有已出和未出的付费课程和项目</li><li>终身会员服务，终身答疑指导</li><li>一对一模拟面试</li><li>等等</li></ul>\n</li><li><strong>go开发Offer收割社群</strong> <ul>\n<li>Go全套资料，包含学习，面试真题等等</li><li>一对一模拟面试，多次</li><li>模拟面试通过可获内推，跟踪整个流程</li><li>终身服务</li><li>详情介绍请看<a href=\"/pay/go.md\">go开发offer收割社群详细介绍</a>了解</li></ul>\n</li></ul>\n<blockquote>\n<p>如有意向，可联系我了解详情</p>\n</blockquote>\n<p><img src=\"./img/image-20220215214357342.png\" alt=\"image-20220215214357342\"></p>\n",
	//		UserName: "悠米",
	//	},
	//}})
	//db.AutoMigrate(&models.User{})
}
