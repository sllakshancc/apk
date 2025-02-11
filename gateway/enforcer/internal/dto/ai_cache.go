package dto

// LLMResponse defines llm response structure
type LLMResponse struct {
	JSON struct {
		Value string `json:"value"`
	} `json:"json"`
}
