// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

import "github.com/gogf/gf/v2/os/gtime"

type CommentSearch struct {
	UserId int64 `json:"user_id" description:"用户id" `

	UserComment string `json:"user_comment" description:"评论内容" `

	CommentTime []string `json:"comment_time" description:"评论时间" `

	Status int `json:"status" description:"数据状态" `

	BookId int64 `json:"book_id" description:"书籍id" `
}

type CommentSave struct {
	UserId int64 `json:"user_id"  description:"用户id" `

	UserComment string `json:"user_comment"  description:"评论内容" `

	CommentTime *gtime.Time `json:"comment_time"  description:"评论时间" `

	Status int `json:"status"  description:"数据状态" `

	BookId int64 `json:"book_id"  description:"书籍id" `
}

type CommentUpdate struct {
	Id int64 `json:"id"  description:"id" `

	UserId int64 `json:"user_id"  description:"用户id" `

	UserComment string `json:"user_comment"  description:"评论内容" `

	CommentTime *gtime.Time `json:"comment_time"  description:"评论时间" `

	Status int `json:"status"  description:"数据状态" `

	BookId int64 `json:"book_id"  description:"书籍id" `
}
