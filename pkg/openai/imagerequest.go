package openai

type ImageRequest struct {
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `json:"n"`
	// The size of the generated images.
	// Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size"`
	// The format in which the generated images are returned.
	// Must be one of url or b64_json
	ResponseFormat string `json:"response_format"`
}