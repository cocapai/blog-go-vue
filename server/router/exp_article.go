package router

import (
	"gin-vue-admin/middleware"
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("article").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("article", v1.CreateExaArticle)     // 创建客户
		ApiRouter.PUT("article", v1.UpdateExaArticle)      // 更新客户
		ApiRouter.DELETE("article", v1.DeleteExaArticle)   // 删除客户
		ApiRouter.GET("article", v1.GetExaArticle)         // 获取单一客户信息
		ApiRouter.GET("articleList", v1.GetExaArticleList) // 获取客户列表
	}
}
