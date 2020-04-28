package model

import (
	"github.com/jinzhu/gorm"
)

type ExaArticle struct {
	gorm.Model
	ArticleName        string  `json:"ArticleName" form:"ArticleName"`
	Content            string  `json:"Content" form:"Content"`
	ArticlePhoneData   string  `json:"ArticlePhoneData" form:"ArticlePhoneData"`
	SysUserID          uint    `json:"sysUserId" form:"sysUserId"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID"`
	SysUser            SysUser `json:"sysUser" form:"sysUser"`
}
