package dto

type GptRequest struct {
	ModelUri          string             `json:"modelUri"`
	CompletionOptions *CompletionOptions `json:"completionOptions"`
	Messages          []*GptMessage      `json:"messages"`
}

type CompletionOptions struct {
	Stream           bool              `json:"stream"`
	Temperature      float64           `json:"temperature"`
	MaxTokens        int               `json:"maxTokens"`
	ReasoningOptions *ReasoningOptions `json:"reasoningOptions"`
}

type ReasoningOptions struct {
	Mode string `json:"mode"`
}

type GptMessage struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type Alternative struct {
	GptMessage *GptMessage `json:"message"`
	Status     string      `json:"status"`
}

type GptResponse struct {
	Result *GptResult `json:"result"`
}

type GptResult struct {
	Alternatives []*Alternative `json:"alternatives"`
	Usage        *Usage         `json:"usage"`
	ModelVersion string         `json:"modelVersion"`
}

type Usage struct {
	InputTextTokens  string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}
