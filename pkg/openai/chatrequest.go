package openai

type ChatRequest struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}
type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
