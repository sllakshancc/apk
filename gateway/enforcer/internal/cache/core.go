package cache

import (
	"encoding/json"
	"fmt"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_service_proc_v3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/wso2/apk/gateway/enforcer/internal/datastore"
	"github.com/wso2/apk/gateway/enforcer/internal/dto"
)

const threshold float32 = 0.8

// CheckCacheForKey checks if the key is in the cache
func CheckCacheForKey(key string, cacheStore datastore.CacheStore, vectorStore VectorProvider, embeddingProvider EmbeddingProvider) (string, error) {

	var response string
	var err error

	response, err = cacheStore.Get(key)
	if err != nil {
		// KV cache check fails. perform similarity search
		response, err = performSimilaritySearch(key, vectorStore, embeddingProvider)
	}

	return response, err
}

// Caches the response value
func cacheResponse(key string, value string, cacheStore datastore.CacheStore) {
	err := cacheStore.Set(key, value)
	if err != nil {
		fmt.Printf("[AI-CACHE] cache set failed, key: %s, error: %v", key, err)
		return
	}
	fmt.Printf("[AI-CACHE] cache set success, key: %s, length of value: %d", key, len(value))

}

func performSimilaritySearch(key string, vectorStore VectorProvider, embeddingProvider EmbeddingProvider) (string, error) {

	if vectorStore == nil {
		return "", fmt.Errorf("performSimilaritySearch fails. vector store not initialized")
	}

	queryEmd, err := embeddingProvider.GetEmbedding(key)
	if err != nil {
		return "", fmt.Errorf("performSimilaritySearch fails. error: %v", err)
	}

	queryResult, err := vectorStore.QueryEmbedding(float64ToFloat32(queryEmd))
	if err != nil {
		return "", fmt.Errorf("performSimilaritySearch fails. error: %v", err)
	}
	if len(queryResult) == 0 {
		return "", fmt.Errorf("vector query results is empty")
	}

	mostSimilarQuery := queryResult[0]
	fmt.Printf("[handleQueryResults] for key: %s, the most similar key found: %s with score: %f", key, mostSimilarQuery.Text, mostSimilarQuery.Score)
	// check threshold
	if mostSimilarQuery.Score < threshold {
		return "", fmt.Errorf("no query above threshold")
	}

	// TODO: make sure answer is not empty
	return mostSimilarQuery.Answer, nil
}

func uploadEmbeddingAndAnswer(key string, value string, vectorStore VectorProvider, embeddingProvider EmbeddingProvider) {

	if vectorStore == nil {
		fmt.Printf("uploadEmbeddingAndAnswer fails. vector store not initialized")
		return
	}

	queryEmd, err := embeddingProvider.GetEmbedding(key)
	if err != nil {
		fmt.Printf("genarating embedding for uploadEmbeddingAndAnswer fails. error: %v", err)
		return
	}

	err = vectorStore.UploadAnswerAndEmbedding(key, float64ToFloat32(queryEmd), value)
	if err != nil {
		fmt.Printf("UploadAnswerAndEmbedding fails. error: %v", err)
		return
	}
}

// SendCachedHTTPResponse makes ext_proc response for cached value
func SendCachedHTTPResponse(cachedResponse string, resp *envoy_service_proc_v3.ProcessingResponse) {

	llmResponse := dto.LLMResponse{
		ID:      "chatcmpl-123", // You may want to generate a unique ID
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   "gpt-3.5-turbo",
		Usage: dto.Usage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
		Choices: []dto.Choice{
			{
				Index: 0,
				Message: dto.Message{
					Role:    "assistant",
					Content: cachedResponse,
				},
				Delta:        []any{nil},
				FinishReason: "stop",
			},
		},
	}

	httpBody, _ := json.Marshal(llmResponse)
	httpBodyLength := len(httpBody)

	headers := &envoy_service_proc_v3.HeaderMutation{
		SetHeaders: []*corev3.HeaderValueOption{
			{
				Header: &corev3.HeaderValue{
					Key:      "Content-Length",
					RawValue: []byte(fmt.Sprintf("%d", httpBodyLength)),
				},
			},
			{
				Header: &corev3.HeaderValue{
					Key:      "Content-Type",
					RawValue: []byte("application/json"),
				},
			},
			{
				Header: &corev3.HeaderValue{
					Key:      "X-Cache-Status",
					RawValue: []byte("HIT"),
				},
			},
		},
	}

	rbq := &envoy_service_proc_v3.ImmediateResponse{
		Status: &v32.HttpStatus{
			Code: v32.StatusCode_OK,
		},
		Headers: headers,
		Body:    httpBody,
	}

	resp.Response = &envoy_service_proc_v3.ProcessingResponse_ImmediateResponse{
		ImmediateResponse: rbq,
	}

}

// TODO: should handled in specific vector provider
func float64ToFloat32(arr []float64) []float32 {
	out := make([]float32, len(arr))
	for i, v := range arr {
		out[i] = float32(v)
	}
	return out
}
