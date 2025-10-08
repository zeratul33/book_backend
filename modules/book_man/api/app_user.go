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

type IndexAppUserReq struct {
	g.Meta `path:"/appUser/index" method:"get" tags:"用户" summary:"分页列表" x-permission:"book_man:appUser:index" `
	model.AuthorHeader
	model.PageListReq
	req.AppUserSearch
}

type IndexAppUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AppUser `json:"items"  dc:"list" `
}

type ListAppUserReq struct {
	g.Meta `path:"/appUser/list" method:"get" tags:"用户" summary:"列表" x-permission:"book_man:appUser:list" `
	model.AuthorHeader
	model.ListReq
	req.AppUserSearch
}

type ListAppUserRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.AppUser `json:"data"  dc:"list" `
}

type SaveAppUserReq struct {
	g.Meta `path:"/appUser/save" method:"post" tags:"用户" summary:"新增" x-permission:"book_man:appUser:save"`
	model.AuthorHeader
	req.AppUserSave
}

type SaveAppUserRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadAppUserReq struct {
	g.Meta `path:"/appUser/read/{Id}" method:"get" tags:"用户" summary:"获取单个信息" x-permission:"book_man:appUser:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"用户 id" v:"required|min:1#Id不能为空"`
}

type ReadAppUserRes struct {
	g.Meta `mime:"application/json"`
	Data   res.AppUser `json:"data" dc:"信息数据"`
}

type UpdateAppUserReq struct {
	g.Meta `path:"/appUser/update/{Id}" method:"put" tags:"用户" summary:"更新" x-permission:"book_man:appUser:update"`
	model.AuthorHeader
	req.AppUserUpdate
}

type UpdateAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteAppUserReq struct {
	g.Meta `path:"/appUser/delete" method:"delete" tags:"用户" summary:"删除" x-permission:"book_man:appUser:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleAppUserReq struct {
	g.Meta `path:"/appUser/recycle" method:"get" tags:"用户" summary:"回收站列表" x-permission:"book_man:appUser:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.AppUserSearch
}

type RecycleAppUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AppUser `json:"items"  dc:"list" `
}

type RealDeleteAppUserReq struct {
	g.Meta `path:"/appUser/realDelete" method:"delete" tags:"用户" summary:"单个或批量真实删除 （清空回收站）" x-permission:"book_man:appUser:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryAppUserReq struct {
	g.Meta `path:"/appUser/recovery" method:"put" tags:"用户" summary:"单个或批量恢复在回收站的" x-permission:"book_man:appUser:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type ExportAppUserReq struct {
	g.Meta `path:"/appUser/export" method:"post" tags:"用户" summary:"导出" x-permission:"book_man:appUser:export"`
	model.AuthorHeader
	model.ListReq
	req.AppUserSearch
}

type ExportAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type ImportAppUserReq struct {
	g.Meta `path:"/appUser/import" method:"post" mime:"multipart/form-data" tags:"用户" summary:"导入" x-permission:"book_man:appUser:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateAppUserReq struct {
	g.Meta `path:"/appUser/downloadTemplate" method:"post,get" tags:"用户" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"book_man:appUser:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusAppUserReq struct {
	g.Meta `path:"/appUser/changeStatus" method:"put" tags:"用户" summary:"更改状态" x-permission:"book_man:appUser:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationAppUserReq struct {
	g.Meta `path:"/appUser/numberOperation" method:"put" tags:"用户" summary:"数字运算操作" x-permission:"book_man:appUser:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"0" v:"min:0#数字不能为空"`
}

type NumberOperationAppUserRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteAppUserReq struct {
	g.Meta `path:"/appUser/remote" method:"post" tags:"用户" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"book_man:appUser:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteAppUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.AppUser `json:"items"  dc:"list" `
	Data  []res.AppUser `json:"data"  dc:"list" `
}
