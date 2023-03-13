package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

var history map[int64][]Messages = make(map[int64][]Messages)
var historyDepth int = 8

func SendRequest(userId int64, prompt string) *ChatResponse {
	url := "https://api.openai.com/v1/chat/completions"
	reqData := prepareRequest(userId, prompt)
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	token := *cli.FlagOpenAIToken
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer " + token)
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
	} else if len(userHist) >= historyDepth {
		userHist = userHist[len(userHist) - historyDepth:]
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
