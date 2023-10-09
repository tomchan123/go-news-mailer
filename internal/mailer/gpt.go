package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GPTServer struct {
	ApiKey      string
	ApiEndpoint string
	ApiModel    string
}

type ChatRepsonse struct {
	Choices []ChatChoice `json:"choices"`
}

type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
}

type ChatMessage struct {
	Content string `json:"content"`
}

func CreateGPTServer() *GPTServer {
	return &GPTServer{
		ApiKey:      "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",
		ApiEndpoint: "https://api.openai.com/v1",
		ApiModel:    "gpt-3.5-turbo",
	}
}

func (gpts *GPTServer) ChatSummarise(txt string) (string, error) {
	url := fmt.Sprintf("%s/chat/completions", gpts.ApiEndpoint)

	return gpts.getChatSummariseApi(url, txt)
}

func (gpts *GPTServer) getChatSummariseApi(url string, txt string) (string, error) {
	rb, err := json.Marshal(map[string]interface{}{
		"model": gpts.ApiModel,
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": `Summarise the following news article you are provided with for a normal adult using no more than 60 words`,
			},
			{
				"role":    "user",
				"content": txt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("GPTServer.getChatSummariseApi: %v", err)
	}

	rbt := bytes.NewBuffer(rb)
	req, err := http.NewRequest(http.MethodPost, url, rbt)
	if err != nil {
		return "", fmt.Errorf("GPTServer.getChatSummariseApi: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gpts.ApiKey))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("GPTServer.getChatSummariseApi: %v", err)
	}

	rpb, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("NewsServer.getArticlesAPIRequest: %v", err)
	}

	resp := ChatRepsonse{}
	if err := json.Unmarshal(rpb, &resp); err != nil {
		panic(err)
	}

	return gpts.getChatcontent(&resp), nil
}

func (gpts *GPTServer) getChatcontent(cr *ChatRepsonse) string {
	// always pick the first choice
	for _, v := range cr.Choices {
		if v.Index == 0 {
			return v.Message.Content
		}
	}

	return ""
}
