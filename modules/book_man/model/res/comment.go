// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"devinggo/internal/model/entity"
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/os/gtime"
)

type Comment struct {
	Id int64 `json:"id"  description:"id" `

	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间" `

	UpdatedAt *gtime.Time `json:"updated_at"  description:"修改时间" `

	DeletedAt *gtime.Time `json:"deleted_at"  description:"软删除" `

	CreatedBy int64 `json:"created_by"  description:"创建者" `

	CreatedByRelate model.UserRelate `json:"created_by_related"  description:"创建人关联信息" `

	UpdatedBy int64 `json:"updated_by"  description:"修改者" `

	UserId int64 `json:"user_id"  description:"用户id" `

	UserComment string `json:"user_comment"  description:"评论内容" `

	CommentTime *gtime.Time `json:"comment_time"  description:"评论时间" `

	Status int `json:"status"  description:"数据状态" `

	BookId int64 `json:"book_id"  description:"书籍id" `
}

type CommentExcel struct {
	UserId int64 `json:"user_id"  description:"用户id"  excelName:"用户id" excelIndex:"6"  `

	UserComment string `json:"user_comment"  description:"评论内容"  excelName:"评论内容" excelIndex:"7"  `

	CommentTime []string `json:"comment_time"  description:"评论时间"  excelName:"评论时间" excelIndex:"8"  `

	Status int `json:"status"  description:"数据状态"  excelName:"数据状态" excelIndex:"9"  `

	BookId int64 `json:"book_id"  description:"书籍id"  excelName:"书籍id" excelIndex:"10"  `
}
type CommentApp struct {
	Comment *entity.Comment `json:"comment"`
	AppUser *entity.AppUser `json:"user"`
	Book    *entity.Book    `json:"book"`
}
