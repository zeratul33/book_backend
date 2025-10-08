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

type IndexCategoryReq struct {
	g.Meta `path:"/category/index" method:"get" tags:"分类" summary:"分页列表" x-permission:"book_man:category:index" `
	model.AuthorHeader
	model.PageListReq
	req.CategorySearch
}

type IndexCategoryRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Category `json:"items"  dc:"list" `
}

type ListCategoryReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"分类" summary:"列表" x-permission:"book_man:category:list" `
	model.AuthorHeader
	model.ListReq
	req.CategorySearch
}

type ListCategoryRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.Category `json:"data"  dc:"list" `
}

type SaveCategoryReq struct {
	g.Meta `path:"/category/save" method:"post" tags:"分类" summary:"新增" x-permission:"book_man:category:save"`
	model.AuthorHeader
	req.CategorySave
}

type SaveCategoryRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadCategoryReq struct {
	g.Meta `path:"/category/read/{Id}" method:"get" tags:"分类" summary:"获取单个信息" x-permission:"book_man:category:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"分类 id" v:"required|min:1#Id不能为空"`
}

type ReadCategoryRes struct {
	g.Meta `mime:"application/json"`
	Data   res.Category `json:"data" dc:"信息数据"`
}

type UpdateCategoryReq struct {
	g.Meta `path:"/category/update/{Id}" method:"put" tags:"分类" summary:"更新" x-permission:"book_man:category:update"`
	model.AuthorHeader
	req.CategoryUpdate
}

type UpdateCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteCategoryReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"分类" summary:"删除" x-permission:"book_man:category:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleCategoryReq struct {
	g.Meta `path:"/category/recycle" method:"get" tags:"分类" summary:"回收站列表" x-permission:"book_man:category:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.CategorySearch
}

type RecycleCategoryRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Category `json:"items"  dc:"list" `
}

type RealDeleteCategoryReq struct {
	g.Meta `path:"/category/realDelete" method:"delete" tags:"分类" summary:"单个或批量真实删除 （清空回收站）" x-permission:"book_man:category:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryCategoryReq struct {
	g.Meta `path:"/category/recovery" method:"put" tags:"分类" summary:"单个或批量恢复在回收站的" x-permission:"book_man:category:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type ExportCategoryReq struct {
	g.Meta `path:"/category/export" method:"post" tags:"分类" summary:"导出" x-permission:"book_man:category:export"`
	model.AuthorHeader
	model.ListReq
	req.CategorySearch
}

type ExportCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type ImportCategoryReq struct {
	g.Meta `path:"/category/import" method:"post" mime:"multipart/form-data" tags:"分类" summary:"导入" x-permission:"book_man:category:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateCategoryReq struct {
	g.Meta `path:"/category/downloadTemplate" method:"post,get" tags:"分类" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"book_man:category:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusCategoryReq struct {
	g.Meta `path:"/category/changeStatus" method:"put" tags:"分类" summary:"更改状态" x-permission:"book_man:category:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationCategoryReq struct {
	g.Meta `path:"/category/numberOperation" method:"put" tags:"分类" summary:"数字运算操作" x-permission:"book_man:category:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"0" v:"min:0#数字不能为空"`
}

type NumberOperationCategoryRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteCategoryReq struct {
	g.Meta `path:"/category/remote" method:"post" tags:"分类" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"book_man:category:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteCategoryRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.Category `json:"items"  dc:"list" `
	Data  []res.Category `json:"data"  dc:"list" `
}
