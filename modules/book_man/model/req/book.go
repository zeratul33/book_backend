// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type BookSearch struct {
	BookName string `json:"book_name" description:"书籍名称" `

	AuthorName string `json:"author_name" description:"作者" `

	CategoryId int64 `json:"category_id" description:"分类" `
}

type BookSave struct {
	BookName string `json:"book_name"  v:"required"  description:"书籍名称" `

	AuthorName string `json:"author_name"  v:"required"  description:"作者" `

	TagText string `json:"tag_text"  description:"标签文字" `

	CoverUrl string `json:"cover_url"  v:"required"  description:"封面" `

	TagColor string `json:"tag_color"  description:"标签颜色" `

	Rate float32 `json:"rate"  description:"评分" `

	CategoryId int64 `json:"category_id"  v:"required"  description:"分类" `

	Description string `json:"description"  v:"required"  description:"描述" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}

type BookUpdate struct {
	Id int64 `json:"id"  description:"id" `

	BookName string `json:"book_name"  v:"required"  description:"书籍名称" `

	AuthorName string `json:"author_name"  v:"required"  description:"作者" `

	TagText string `json:"tag_text"  description:"标签文字" `

	CoverUrl string `json:"cover_url"  v:"required"  description:"封面" `

	TagColor string `json:"tag_color"  description:"标签颜色" `

	Rate float32 `json:"rate"  description:"评分" `

	CategoryId int64 `json:"category_id"  v:"required"  description:"分类" `

	Description string `json:"description"  v:"required"  description:"描述" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}
