import service from '@/utils/request'

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbmodel.ExaArticle true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [post]
export const createExaArticle = (data) => {
    return service({
        url: "/article/article",
        method: 'post',
        data
    })
}



// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbmodel.ExaArticle true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [put]
export const updateExaArticle = (data) => {
    return service({
        url: "/article/article",
        method: 'put',
        data
    })
}


// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbmodel.ExaArticle true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [delete]
export const deleteExaArticle = (data) => {
    return service({
        url: "/article/article",
        method: 'delete',
        data
    })
}


// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbmodel.ExaArticle true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/Article [get]
export const getExaArticle = (params) => {
    return service({
        url: "/article/article",
        method: 'get',
        params
    })
}


// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/ArticleList [get]
export const getExaArticleList = (params) => {
    return service({
        url: "/article/articleList",
        method: 'get',
        params
    })
}