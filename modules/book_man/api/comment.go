// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package api

import (
	"devinggo/modules/book_man/model/req"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IndexCommentReq struct {
	g.Meta `path:"/comment/index" method:"get" tags:"评论" summary:"分页列表" x-permission:"book_man:comment:index" `
	model.AuthorHeader
	model.PageListReq
	req.CommentSearch
}

type IndexCommentRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Comment `json:"items"  dc:"list" `
}

type ListCommentReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"评论" summary:"列表" x-permission:"book_man:comment:list" `
	model.AuthorHeader
	model.ListReq
	req.CommentSearch
}

type ListCommentRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.Comment `json:"data"  dc:"list" `
}

type SaveCommentReq struct {
	g.Meta `path:"/comment/save" method:"post" tags:"评论" summary:"新增" x-permission:"book_man:comment:save"`
	model.AuthorHeader
	req.CommentSave
}

type SaveCommentRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadCommentReq struct {
	g.Meta `path:"/comment/read/{Id}" method:"get" tags:"评论" summary:"获取单个信息" x-permission:"book_man:comment:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"评论 id" v:"required|min:1#Id不能为空"`
}

type ReadCommentRes struct {
	g.Meta `mime:"application/json"`
	Data   res.Comment `json:"data" dc:"信息数据"`
}

type UpdateCommentReq struct {
	g.Meta `path:"/comment/update/{Id}" method:"put" tags:"评论" summary:"更新" x-permission:"book_man:comment:update"`
	model.AuthorHeader
	req.CommentUpdate
}

type UpdateCommentRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/comment/delete" method:"delete" tags:"评论" summary:"删除" x-permission:"book_man:comment:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCommentRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleCommentReq struct {
	g.Meta `path:"/comment/recycle" method:"get" tags:"评论" summary:"回收站列表" x-permission:"book_man:comment:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.CommentSearch
}

type RecycleCommentRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Comment `json:"items"  dc:"list" `
}

type RealDeleteCommentReq struct {
	g.Meta `path:"/comment/realDelete" method:"delete" tags:"评论" summary:"单个或批量真实删除 （清空回收站）" x-permission:"book_man:comment:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteCommentRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryCommentReq struct {
	g.Meta `path:"/comment/recovery" method:"put" tags:"评论" summary:"单个或批量恢复在回收站的" x-permission:"book_man:comment:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryCommentRes struct {
	g.Meta `mime:"application/json"`
}

type ExportCommentReq struct {
	g.Meta `path:"/comment/export" method:"post" tags:"评论" summary:"导出" x-permission:"book_man:comment:export"`
	model.AuthorHeader
	model.ListReq
	req.CommentSearch
}

type ExportCommentRes struct {
	g.Meta `mime:"application/json"`
}

type ImportCommentReq struct {
	g.Meta `path:"/comment/import" method:"post" mime:"multipart/form-data" tags:"评论" summary:"导入" x-permission:"book_man:comment:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportCommentRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateCommentReq struct {
	g.Meta `path:"/comment/downloadTemplate" method:"post,get" tags:"评论" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"book_man:comment:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateCommentRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusCommentReq struct {
	g.Meta `path:"/comment/changeStatus" method:"put" tags:"评论" summary:"更改状态" x-permission:"book_man:comment:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusCommentRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationCommentReq struct {
	g.Meta `path:"/comment/numberOperation" method:"put" tags:"评论" summary:"数字运算操作" x-permission:"book_man:comment:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"0" v:"min:0#数字不能为空"`
}

type NumberOperationCommentRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteCommentReq struct {
	g.Meta `path:"/comment/remote" method:"post" tags:"评论" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"book_man:comment:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteCommentRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Comment `json:"items"  dc:"list" `
	Data  []res.Comment `json:"data"  dc:"list" `
}
