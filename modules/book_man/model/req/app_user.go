// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type AppUserSearch struct {
	Id int64 `json:"id" description:"id" `

	CreatedAt []string `json:"created_at" description:"创建时间" `

	Username string `json:"username" description:"用户名" `

	Status int `json:"status" description:"数据状态" `
}

type AppUserSave struct {
	Username string `json:"username"  v:"required"  description:"用户名" `

	PasswordHash string `json:"password_hash"  description:"哈希加密密码" `

	Nickname string `json:"nickname"  v:"required"  description:"昵称" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}

type AppUserUpdate struct {
	Id int64 `json:"id"  description:"id" `

	Username string `json:"username"  v:"required"  description:"用户名" `

	PasswordHash string `json:"password_hash"  description:"哈希加密密码" `

	Nickname string `json:"nickname"  v:"required"  description:"昵称" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}
