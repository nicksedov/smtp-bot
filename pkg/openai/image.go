package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nicksedov/sbconn-bot/pkg/cli"
)

func SendImageRequest(prompt string) *ImageResponse {
	url := "https://api.openai.com/v1/images/generations"
	reqData := prepareImageRequest(prompt)
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	token := *cli.FlagOpenAIToken
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer " + token)
	client := GetClient()
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	resp := &ImageResponse{HttpStatus: response.StatusCode}
	error = json.Unmarshal(body, resp)
	if error != nil {
		fmt.Println(err)
		return resp
	}
	return resp
}

func prepareImageRequest (prompt string) *ImageRequest {
	return &ImageRequest{
		Model: "dall-e-3",
		Prompt: prompt,
		N:    1,
		Size: "1024x1024",
		ResponseFormat: "url",
	}
}

