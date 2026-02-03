package cache

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// EmbeddingProvider interface
type EmbeddingProvider interface {
	GetEmbedding(queryString string) ([]float64, error)
}

// SBERT API Configuration
const (
	sbertEndpoint = "http://host.docker.internal:8000/v1/embeddings"
	modelName     = "sentence-transformers/all-MiniLM-L6-v2"
	timeout       = 10 * time.Second
)

// SBERTRequest represents the structure for sending requests to the SBERT API.
type SBERTRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// SBERTResponse represents the response structure from the SBERT API.
type SBERTResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

// SBERTProvider struct
type SBERTProvider struct {
	client *http.Client
}

// CreateNewSBERTProvider initializes and returns a new SBERTProvider.
func CreateNewSBERTProvider() EmbeddingProvider {
	return &SBERTProvider{
		//TODO: use http wrapper for simplify, so host, port can also set here
		client: &http.Client{Timeout: timeout},
	}
}

// Construct request parameters
func (s *SBERTProvider) constructParameters(text string) (string, []byte, error) {
	if text == "" {
		return "", nil, errors.New("queryString cannot be empty")
	}

	requestData := SBERTRequest{
		Input: text,
		Model: modelName,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Failed to marshal request: %v\n", err)
		return "", nil, err
	}

	return sbertEndpoint, requestBody, nil
}

// Parse response and return full SBERTResponse
func (s *SBERTProvider) parseTextEmbedding(resp *http.Response) (*SBERTResponse, error) {
	var sbertResponse SBERTResponse
	err := json.NewDecoder(resp.Body).Decode(&sbertResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(sbertResponse.Data) == 0 || len(sbertResponse.Data[0].Embedding) == 0 {
		return nil, errors.New("no embedding found in response")
	}

	return &sbertResponse, nil
}

// GetEmbedding function (returns only the embedding)
func (s *SBERTProvider) GetEmbedding(queryString string) ([]float64, error) {
	// Construct parameters
	embURL, embRequestBody, err := s.constructParameters(queryString)
	if err != nil {
		return nil, err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", embURL, bytes.NewBuffer(embRequestBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get embedding, status code: %d", resp.StatusCode)
	}

	// Parse response
	sbertResponse, err := s.parseTextEmbedding(resp)
	if err != nil {
		return nil, err
	}

	// Return only the embedding
	return sbertResponse.Data[0].Embedding, nil
}
