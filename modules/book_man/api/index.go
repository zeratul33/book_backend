package api

import (
	"devinggo/modules/book_man/model/req"
	"devinggo/modules/book_man/model/res"
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexPageReq struct {
	g.Meta `path:"/index" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-exceptLogin:"true"`
}

type IndexPageRes struct {
	g.Meta    `mime:"application/json"`
	BookItems []res.Book `json:"bookItems"`
}

type LoginReq struct {
	g.Meta        `path:"/login" method:"post" tags:"IndexApi" x-exceptAuth:"true" x-exceptLogin:"true"`
	req.LoginBody `json:"loginBody"`
}

type LoginRes struct {
	g.Meta    `mime:"application/json"`
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expiresIn"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
}

type RegisterReq struct {
	g.Meta           `path:"/register" method:"post" tags:"IndexApi" x-exceptAuth:"true" x-exceptLogin:"true"`
	req.RegisterBody `json:"registerBody"`
}

type RegisterRes struct {
	g.Meta `mime:"application/json"`
	Result bool `json:"result"`
}

type GetUserInfoReq struct {
	g.Meta `path:"/getUserInfo" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}
type GetUserInfoRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo res.AppUser `json:"userInfo"`
}

type GetCategoryListReq struct {
	g.Meta `path:"/getCategoryList" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}
type GetCategoryListRes struct {
	g.Meta    `mime:"application/json"`
	Categorys []res.Category `json:"categorys"`
}

type GetBookListReq struct {
	g.Meta `path:"/getBookList" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}

type GetBookListRes struct {
	g.Meta `mime:"application/json"`
	Books  []res.Book `json:"books"`
}
type GetSubscribedListReq struct {
	g.Meta `path:"/getSubscribedList" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}
type GetSubscribedListRes struct {
	g.Meta         `mime:"application/json"`
	SubscribedList []res.Subscribed `json:"subscribedList"`
}

type GetCommentByBookReq struct {
	g.Meta `path:"/getCommentByBook/{id}" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}
type GetCommentByBookRes struct {
	g.Meta   `mime:"application/json"`
	Comments []res.CommentApp `json:"comments"`
}

type GetCommentListByUserReq struct {
	g.Meta `path:"/getCommentListByUser" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}

type GetCommentListByUserRes struct {
	g.Meta   `mime:"application/json"`
	Comments []res.CommentApp `json:"comments"`
}

type GetBookByIdReq struct {
	g.Meta `path:"/getBookById/{id}" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}

type GetBookByIdRes struct {
	g.Meta `mime:"application/json"`
	Book   res.Book `json:"book"`
}

type SubscribeBookReq struct {
	g.Meta `path:"/subscribeBook" method:"post" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
	Id int64 `json:"id" v:"required#id不能为空"`
}

type SubscribeBookRes struct {
	g.Meta `mime:"application/json"`
	Result bool `json:"result"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
}
type LogoutRes struct {
	g.Meta `mime:"application/json"`
	Result bool `json:"result"`
}

type StartScriptReq struct {
	g.Meta `path:"/startScript" method:"get" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test" x-exceptLogin:"true"`
	model.AuthorHeader
}
type StartScriptRes struct {
	g.Meta `mime:"application/json"`
	Result bool `json:"result"`
}

type PublishCommentReq struct {
	g.Meta `path:"/publishComment" method:"post" tags:"IndexApi" x-exceptAuth:"true" x-permission:"api:test"`
	model.AuthorHeader
	Comment string `json:"comment" v:"required#评论内容不能为空"`
	BookId  int64  `json:"bookId" v:"required#书籍id不能为空"`
}

type PublishCommentRes struct {
	g.Meta `mime:"application/json"`
	Result bool `json:"result"`
}
