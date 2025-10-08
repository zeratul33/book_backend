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

type Category struct {
	Id int64 `json:"id"  description:"id" `

	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间" `

	UpdatedAt *gtime.Time `json:"updated_at"  description:"修改时间" `

	DeletedAt *gtime.Time `json:"deleted_at"  description:"软删除" `

	CreatedBy int64 `json:"created_by"  description:"创建者" `

	CreatedByRelate model.UserRelate `json:"created_by_related"  description:"创建人关联信息" `

	UpdatedBy int64 `json:"updated_by"  description:"修改者" `

	UpdatedByRelate model.UserRelate `json:"updated_by_related"  description:"更新人关联信息" `

	CategoryName string `json:"category_name"  description:"分类名称" `

	Status int `json:"status"  description:"数据状态" `
}

type CategoryExcel struct {
	CategoryName string `json:"category_name"  v:"required"  description:"分类名称"  excelName:"分类名称" excelIndex:"6"  `

	Status int `json:"status"  v:"required"  description:"数据状态"  excelName:"数据状态" excelIndex:"7"  `
}
