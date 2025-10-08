package controller

import (
	"context"
	"devinggo/modules/book_man/api"
	"devinggo/modules/book_man/global"
	"devinggo/modules/system/controller/base"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	ChatController = chatController{}
)

type chatController struct {
	base.BaseController
	//sk-48af12e0b70e483e89a456799a653a34

}

func (c *chatController) Chat(ctx context.Context, in *api.ChatReq) (out *api.ChatRes, err error) {
	req := g.RequestFromCtx(ctx).Response
	req.Header().Set("Content-Type", "text/event-stream")
	req.Header().Set("Cache-Control", "no-cache")
	req.Header().Set("Connection", "keep-alive")
	req.Header().Set("Access-Control-Allow-Origin", "*")
	messages, err := global.GlobalChatMemory.ChatMemory.ChatHistory.Messages(ctx)

	llm, err := openai.New(
		openai.WithModel("deepseek-chat"),
		openai.WithToken("sk-48af12e0b70e483e89a456799a653a34"),
		openai.WithBaseURL("https://api.deepseek.com/v1"))
	var conversation string = "所有的回答尽可能简洁，不要用markdown格式，直接用普通文本"
	for _, msg := range messages {
		conversation += msg.GetContent() + "\n"
	}

	// Add current input to the conversation
	fullPrompt := conversation + "Human: " + in.Comment + "\nAssistant:"
	prompt, err := llms.GenerateFromSinglePrompt(ctx,
		llm,
		fullPrompt,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Println("---------test---------")
			fmt.Println(string(chunk))
			//out = &api.ChatRes{
			//	Data: string(chunk),
			//}
			req.Write(string(chunk))
			req.Flush()
			return err
		}),
	)
	global.GlobalChatMemory.ChatMemory.ChatHistory.AddUserMessage(ctx, in.Comment)
	global.GlobalChatMemory.ChatMemory.ChatHistory.AddAIMessage(ctx, prompt)
	return nil, nil
}
