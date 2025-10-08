// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type CategorySearch struct {
	CategoryName string `json:"category_name" description:"分类名称" `

	Status int `json:"status" description:"数据状态" `
}

type CategorySave struct {
	CategoryName string `json:"category_name"  v:"required"  description:"分类名称" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}

type CategoryUpdate struct {
	Id int64 `json:"id"  description:"id" `

	CategoryName string `json:"category_name"  v:"required"  description:"分类名称" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}
