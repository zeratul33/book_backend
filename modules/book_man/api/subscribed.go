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

type IndexSubscribedReq struct {
	g.Meta `path:"/subscribed/index" method:"get" tags:"订阅表" summary:"分页列表" x-permission:"book_man:subscribed:index" `
	model.AuthorHeader
	model.PageListReq
	req.SubscribedSearch
}

type IndexSubscribedRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Subscribed `json:"items"  dc:"list" `
}

type ListSubscribedReq struct {
	g.Meta `path:"/subscribed/list" method:"get" tags:"订阅表" summary:"列表" x-permission:"book_man:subscribed:list" `
	model.AuthorHeader
	model.ListReq
	req.SubscribedSearch
}

type ListSubscribedRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.Subscribed `json:"data"  dc:"list" `
}

type SaveSubscribedReq struct {
	g.Meta `path:"/subscribed/save" method:"post" tags:"订阅表" summary:"新增" x-permission:"book_man:subscribed:save"`
	model.AuthorHeader
	req.SubscribedSave
}

type SaveSubscribedRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadSubscribedReq struct {
	g.Meta `path:"/subscribed/read/{Id}" method:"get" tags:"订阅表" summary:"获取单个信息" x-permission:"book_man:subscribed:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"订阅表 id" v:"required|min:1#Id不能为空"`
}

type ReadSubscribedRes struct {
	g.Meta `mime:"application/json"`
	Data   res.Subscribed `json:"data" dc:"信息数据"`
}

type UpdateSubscribedReq struct {
	g.Meta `path:"/subscribed/update/{Id}" method:"put" tags:"订阅表" summary:"更新" x-permission:"book_man:subscribed:update"`
	model.AuthorHeader
	req.SubscribedUpdate
}

type UpdateSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSubscribedReq struct {
	g.Meta `path:"/subscribed/delete" method:"delete" tags:"订阅表" summary:"删除" x-permission:"book_man:subscribed:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleSubscribedReq struct {
	g.Meta `path:"/subscribed/recycle" method:"get" tags:"订阅表" summary:"回收站列表" x-permission:"book_man:subscribed:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SubscribedSearch
}

type RecycleSubscribedRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Subscribed `json:"items"  dc:"list" `
}

type RealDeleteSubscribedReq struct {
	g.Meta `path:"/subscribed/realDelete" method:"delete" tags:"订阅表" summary:"单个或批量真实删除 （清空回收站）" x-permission:"book_man:subscribed:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type RecoverySubscribedReq struct {
	g.Meta `path:"/subscribed/recovery" method:"put" tags:"订阅表" summary:"单个或批量恢复在回收站的" x-permission:"book_man:subscribed:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoverySubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type ExportSubscribedReq struct {
	g.Meta `path:"/subscribed/export" method:"post" tags:"订阅表" summary:"导出" x-permission:"book_man:subscribed:export"`
	model.AuthorHeader
	model.ListReq
	req.SubscribedSearch
}

type ExportSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type ImportSubscribedReq struct {
	g.Meta `path:"/subscribed/import" method:"post" mime:"multipart/form-data" tags:"订阅表" summary:"导入" x-permission:"book_man:subscribed:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateSubscribedReq struct {
	g.Meta `path:"/subscribed/downloadTemplate" method:"post,get" tags:"订阅表" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"book_man:subscribed:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusSubscribedReq struct {
	g.Meta `path:"/subscribed/changeStatus" method:"put" tags:"订阅表" summary:"更改状态" x-permission:"book_man:subscribed:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationSubscribedReq struct {
	g.Meta `path:"/subscribed/numberOperation" method:"put" tags:"订阅表" summary:"数字运算操作" x-permission:"book_man:subscribed:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"0" v:"min:0#数字不能为空"`
}

type NumberOperationSubscribedRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteSubscribedReq struct {
	g.Meta `path:"/subscribed/remote" method:"post" tags:"订阅表" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"book_man:subscribed:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteSubscribedRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Subscribed `json:"items"  dc:"list" `
	Data  []res.Subscribed `json:"data"  dc:"list" `
}
