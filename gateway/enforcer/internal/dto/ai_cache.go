package dto

// LLMResponse defines llm response structure
type LLMResponse struct {
	JSON struct {
		Value string `json:"value"`
	} `json:"json"`
}

// LLMRequest defines llm Request structure
type LLMRequest struct {
	Key string `json:"key"`
}

// GetValue Get the value from llm response
func (r *LLMResponse) GetValue() (string, bool) {
	if r.JSON.Value == "" {
		return "", false
	}
	return r.JSON.Value, true
}

// GetKey Get the key from llm response
func (r *LLMRequest) GetKey() (string, bool) {
	if r.Key == "" {
		return "", false
	}
	return r.Key, true
}
