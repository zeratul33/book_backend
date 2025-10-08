// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SubscribedSearch struct {
	Id int64 `json:"id" description:"id" `

	SubscribedUser int64 `json:"subscribed_user" description:"订阅用户" `

	SubscriebedBook int64 `json:"subscriebed_book" description:"订阅书籍" `

	Status int `json:"status" description:"数据状态" `
}

type SubscribedSave struct {
	SubscribedUser int64 `json:"subscribed_user"  v:"required"  description:"订阅用户" `

	SubscriebedBook int64 `json:"subscriebed_book"  v:"required"  description:"订阅书籍" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}

type SubscribedUpdate struct {
	Id int64 `json:"id"  description:"id" `

	SubscribedUser int64 `json:"subscribed_user"  v:"required"  description:"订阅用户" `

	SubscriebedBook int64 `json:"subscriebed_book"  v:"required"  description:"订阅书籍" `

	Status int `json:"status"  v:"required"  description:"数据状态" `
}
