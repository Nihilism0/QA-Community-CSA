package api

import (
	"CSAwork/utils/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", register) // 注册
	r.POST("/login", login)       // 登录
	QARouter := r.Group("/qa")    //QA是Question和Answer
	{
		QARouter.Use(middleware.JWTAuthMiddleware()) //我用了一个不知道算不算正规的api名字设计 首字母为q或a表明是问题还是回答 后面跟动词
		QARouter.POST("/qcreate", Qcreate)           //发布问题
		QARouter.GET("/qsubmited", Qsubmited)        //查看某人发表过的问题及其解答
		QARouter.POST("/acreate", Acreate)           //回答问题
		QARouter.GET("/asubmited", Asubmited)        //查看某人回答过的回答及其问题
		QARouter.PUT("/qmodify", Qmodify)            //修改问题	新学一个单词==>modify  v.修改
		QARouter.PUT("/amodify", Amodify)            //修改回答
		QARouter.DELETE("/qdelete", Qdelete)         //删除问题
		QARouter.DELETE("/adelete", Adelete)         //删除回答
		QARouter.POST("/acomment", Acomment)         //评论回答
		//接下来使用redis缓存
		QARouter.POST("/praise", Praise)               //点赞问题
		QARouter.DELETE("/cancelpraise", CancelPraise) //取消点赞
	}
	r.GET("/seepraise", SeePraise) //看问题几个赞
	r.Run(":3920")                 // 跑在 3920 端口上
}
