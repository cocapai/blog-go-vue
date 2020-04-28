package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateExaArticle
// @description   create a Article, 创建用户
// @param     e               model.ExaArticle
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateExaArticle(e model.ExaArticle) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a Article, 删除用户
// @auth                     （2020/04/05  20:22）
// @param     e               *model.ExaArticle
// @return                    error

func DeleteExaArticle(e model.ExaArticle) (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

// @title    UpdateExaArticle
// @description   update a Article, 更新用户
// @param     e               *model.ExaArticle
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateExaArticle(e *model.ExaArticle) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

// @title    GetExaArticle
// @description   get the info of a costumer , 获取用户信息
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Article        ExaArticle

func GetExaArticle(id uint) (err error, Article model.ExaArticle) {
	err = global.GVA_DB.Where("id = ?", id).First(&Article).Error
	return
}

// @title    GetArticleInfoList
// @description   get Article list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     sysUserAuthorityID              string
// @param     info            PageInfo
// @return                    error

func GetArticleInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var a model.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := GetAuthorityInfo(a)
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var ArticleList []model.ExaArticle
	err = db.Where("sys_user_authority_id in (?)", dataId).Find(&ArticleList).Count(&total).Error
	if err != nil {
		return err, ArticleList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in (?)", dataId).Find(&ArticleList).Error
	}
	return err, ArticleList, total
}
