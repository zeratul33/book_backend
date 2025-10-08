// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/os/gtime"
)

type Book struct {
	Id int64 `json:"id"  description:"id" `

	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间" `

	UpdatedAt *gtime.Time `json:"updated_at"  description:"修改时间" `

	DeletedAt *gtime.Time `json:"deleted_at"  description:"软删除" `

	CreatedBy int64 `json:"created_by"  description:"创建者" `

	CreatedByRelate model.UserRelate `json:"created_by_related"  description:"创建人关联信息" `

	UpdatedBy int64 `json:"updated_by"  description:"修改者" `

	UpdatedByRelate model.UserRelate `json:"updated_by_related"  description:"更新人关联信息" `

	BookName string `json:"book_name"  description:"书籍名称" `

	AuthorName string `json:"author_name"  description:"作者" `

	TagText string `json:"tag_text"  description:"标签文字" `

	CoverUrl string `json:"cover_url"  description:"封面" `

	TagColor string `json:"tag_color"  description:"标签颜色" `

	Rate float32 `json:"rate"  description:"评分" `

	CategoryId int64 `json:"category_id"  description:"分类" `

	Description string `json:"description"  description:"描述" `

	Status int `json:"status"  description:"数据状态" `
}

type BookExcel struct {
	BookName string `json:"book_name"  v:"required"  description:"书籍名称"  excelName:"书籍名称" excelIndex:"6"  `

	AuthorName string `json:"author_name"  v:"required"  description:"作者"  excelName:"作者" excelIndex:"7"  `

	TagText string `json:"tag_text"  description:"标签文字"  excelName:"标签文字" excelIndex:"8"  `

	CoverUrl string `json:"cover_url"  v:"required"  description:"封面"  excelName:"封面" excelIndex:"9"  `

	TagColor string `json:"tag_color"  description:"标签颜色"  excelName:"标签颜色" excelIndex:"10"  `

	Rate float32 `json:"rate"  description:"评分"  excelName:"评分" excelIndex:"11"  `

	CategoryId int64 `json:"category_id"  v:"required"  description:"分类"  excelName:"分类" excelIndex:"12"  `

	Description string `json:"description"  v:"required"  description:"描述"  excelName:"描述" excelIndex:"13"  `

	Status int `json:"status"  v:"required"  description:"数据状态"  excelName:"数据状态" excelIndex:"14"  `
}
