package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var history map[int64][]Messages = make(map[int64][]Messages)

func SendRequest(userId int64, content string) *ChatResponse {
	url := "https://api.openai.com/v1/chat/completions"
	reqData := prepareRequest(userId, content)
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer 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")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	resp := &ChatResponse{}
	error = json.Unmarshal(body, resp)
	if error != nil {
		fmt.Println(err)
		return resp
	}
	processResponse(userId, resp)
	return resp
}

func updateHistory(userId int64, role string, content string) {
	userHist := history[userId]
	if userHist == nil {
		userHist = []Messages{}
	}
	userHist = append(userHist, Messages{Role: role, Content: content})
	history[userId] = userHist
}

func prepareRequest(userId int64, cotent string) *ChatRequest {
	updateHistory(userId, "user", cotent)
	req := ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: history[userId],
	}
	return &req
}

func processResponse(userId int64, resp *ChatResponse) {
	choices := resp.Choices
	if len(choices) > 0 {
		msg := choices[0]
		updateHistory(userId, msg.Message.Role, msg.Message.Content)
	}
}
