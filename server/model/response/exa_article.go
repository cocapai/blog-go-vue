package response

import "gin-vue-admin/model"

type ExaArticleResponse struct {
	Article model.ExaArticle `json:"article"`
}
