package main

import (
	"context"
	"fmt"

	"github.com/otiai10/openaigo"
)

// 基本聊天
func baseChat(client *openaigo.Client, ctx context.Context) {
	request := openaigo.CompletionRequestBody{
		Model:  "text-davinci-003",
		Prompt: []string{"请写100个字以上来描述下chatgpt", "你好呀"},
	}
	response, err := client.Completion(ctx, request)
	fmt.Println(response, err)
}

func createImage(client *openaigo.Client, ctx context.Context) {

	response, err := client.CreateImage(ctx, openaigo.ImageGenerationRequestBody{
		Prompt: "a beautiful two-dimensional girl",
		N:      2,
		Size:   "1024x1024",
	})
	fmt.Println(response, err)
}

func main() {
	// client := openaigo.NewClient(os.Getenv("OPENAI_API_KEY"))
	client := openaigo.NewClient("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")

	ctx := context.Background()
	// baseChat(client, ctx)
	createImage(client, ctx)

}
