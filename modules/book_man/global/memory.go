package global

import "github.com/tmc/langchaingo/memory"

var GlobalChatMemory = &Memory{}

type Memory struct {
	ChatMemory *memory.ConversationBuffer
}

func init() {
	GlobalChatMemory = &Memory{
		ChatMemory: memory.NewConversationBuffer(),
	}
}
