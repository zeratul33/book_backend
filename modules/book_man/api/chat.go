package api

import "github.com/gogf/gf/v2/frame/g"

type ChatReq struct {
	g.Meta  `path:"/chat" method:"post" tags:"book_man" summary:"对话" x-exceptAuth:"true" x-exceptLogin:"true" x-permission:"book_man:test"`
	Comment string `json:"comment"`
}

type ChatRes struct {
	g.Meta `mime:"application/octet-stream" `
	Data   string `json:"data"`
}
