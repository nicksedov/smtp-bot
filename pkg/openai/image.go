package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	resp := &ImageResponse{}
	error = json.Unmarshal(body, resp)
	if error != nil {
		fmt.Println(err)
		return resp
	}
	return resp
}

func prepareImageRequest (prompt string) *ImageRequest {
	return &ImageRequest{
		Prompt: prompt,
		N:    1,
		Size: "512x512",
		ResponseFormat: "url",
	}
}

