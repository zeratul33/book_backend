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

type IndexBookReq struct {
	g.Meta `path:"/book/index" method:"get" tags:"书籍" summary:"分页列表" x-permission:"book_man:book:index" `
	model.AuthorHeader
	model.PageListReq
	req.BookSearch
}

type IndexBookRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Book `json:"items"  dc:"list" `
}

type ListBookReq struct {
	g.Meta `path:"/book/list" method:"get" tags:"书籍" summary:"列表" x-permission:"book_man:book:list" `
	model.AuthorHeader
	model.ListReq
	req.BookSearch
}

type ListBookRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.Book `json:"data"  dc:"list" `
}

type SaveBookReq struct {
	g.Meta `path:"/book/save" method:"post" tags:"书籍" summary:"新增" x-permission:"book_man:book:save"`
	model.AuthorHeader
	req.BookSave
}

type SaveBookRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadBookReq struct {
	g.Meta `path:"/book/read/{Id}" method:"get" tags:"书籍" summary:"获取单个信息" x-permission:"book_man:book:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"书籍 id" v:"required|min:1#Id不能为空"`
}

type ReadBookRes struct {
	g.Meta `mime:"application/json"`
	Data   res.Book `json:"data" dc:"信息数据"`
}

type UpdateBookReq struct {
	g.Meta `path:"/book/update/{Id}" method:"put" tags:"书籍" summary:"更新" x-permission:"book_man:book:update"`
	model.AuthorHeader
	req.BookUpdate
}

type UpdateBookRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteBookReq struct {
	g.Meta `path:"/book/delete" method:"delete" tags:"书籍" summary:"删除" x-permission:"book_man:book:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteBookRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleBookReq struct {
	g.Meta `path:"/book/recycle" method:"get" tags:"书籍" summary:"回收站列表" x-permission:"book_man:book:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.BookSearch
}

type RecycleBookRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Book `json:"items"  dc:"list" `
}

type RealDeleteBookReq struct {
	g.Meta `path:"/book/realDelete" method:"delete" tags:"书籍" summary:"单个或批量真实删除 （清空回收站）" x-permission:"book_man:book:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteBookRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryBookReq struct {
	g.Meta `path:"/book/recovery" method:"put" tags:"书籍" summary:"单个或批量恢复在回收站的" x-permission:"book_man:book:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryBookRes struct {
	g.Meta `mime:"application/json"`
}

type ExportBookReq struct {
	g.Meta `path:"/book/export" method:"post" tags:"书籍" summary:"导出" x-permission:"book_man:book:export"`
	model.AuthorHeader
	model.ListReq
	req.BookSearch
}

type ExportBookRes struct {
	g.Meta `mime:"application/json"`
}

type ImportBookReq struct {
	g.Meta `path:"/book/import" method:"post" mime:"multipart/form-data" tags:"书籍" summary:"导入" x-permission:"book_man:book:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportBookRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateBookReq struct {
	g.Meta `path:"/book/downloadTemplate" method:"post,get" tags:"书籍" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"book_man:book:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateBookRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusBookReq struct {
	g.Meta `path:"/book/changeStatus" method:"put" tags:"书籍" summary:"更改状态" x-permission:"book_man:book:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusBookRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationBookReq struct {
	g.Meta `path:"/book/numberOperation" method:"put" tags:"书籍" summary:"数字运算操作" x-permission:"book_man:book:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"0" v:"min:0#数字不能为空"`
}

type NumberOperationBookRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteBookReq struct {
	g.Meta `path:"/book/remote" method:"post" tags:"书籍" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"book_man:book:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteBookRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Book `json:"items"  dc:"list" `
	Data  []res.Book `json:"data"  dc:"list" `
}
