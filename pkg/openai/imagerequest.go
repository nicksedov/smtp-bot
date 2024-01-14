package openai

type ImageRequest struct {
	// A name of the generative model
	Model string `json:"model"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `json:"n"`
	// The size of the generated images.
	// Accepted values depend on a model
	Size string `json:"size"`
	// The format in which the generated images are returned.
	// Must be one of url or b64_json
	ResponseFormat string `json:"response_format"`
}