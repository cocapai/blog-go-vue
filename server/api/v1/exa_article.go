package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaArticle true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [post]
func CreateExaArticle(c *gin.Context) {
	var cu model.ExaArticle
	_ = c.ShouldBindJSON(&cu)
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	cu.SysUserID = waitUse.ID
	cu.SysUserAuthorityID = waitUse.AuthorityId
	err := service.CreateExaArticle(cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaArticle true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [delete]
func DeleteExaArticle(c *gin.Context) {
	var cu model.ExaArticle
	_ = c.ShouldBindJSON(&cu)
	err := service.DeleteExaArticle(cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaArticle true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [put]
func UpdateExaArticle(c *gin.Context) {
	var cu model.ExaArticle
	_ = c.ShouldBindJSON(&cu)
	err := service.UpdateExaArticle(&cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败：%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaArticle true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [get]
func GetExaArticle(c *gin.Context) {
	var cu model.ExaArticle
	_ = c.ShouldBindQuery(&cu)
	err, Article := service.GetExaArticle(cu.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.ExaArticleResponse{Article: Article}, c)
	}
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/ArticleList [get]
func GetExaArticleList(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, ArticleList, total := service.GetArticleInfoList(waitUse.AuthorityId, pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     ArticleList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
