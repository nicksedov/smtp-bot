package openai

type ImageResponse struct {
	HttpStatus int `json:"status"`
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}