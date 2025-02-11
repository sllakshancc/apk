package datastore

import (
	"errors"
	"fmt"
	"sync"
)

// CacheStore defines an interface for a simple key-value cache.
type CacheStore interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

// MockCache is an in-memory implementation of CacheStore for testing purposes.
type MockCache struct {
	store map[string]string
	mu    sync.RWMutex
}

// Set stores a key-value pair in the mock cache.
func (s *MockCache) Set(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
	fmt.Printf("[MockCache] SET key=%q -> value=%q\n", key, value)
	return nil
}

// Get retrieves a value from the mock cache.
// If the key does not exist, it returns an error.
func (s *MockCache) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.store[key]
	fmt.Printf("[MockCache] GET key=%q -> value=%q, exists=%v\n", key, val, exists)
	if !exists {
		return "", errors.New("key doesn't exist")
	}
	return val, nil
}

// NewMockCache initializes and returns a new instance of MockCache.
func NewMockCache() *MockCache {
	return &MockCache{
		store: make(map[string]string),
	}
}

// IncomingRequestCacheKeyStore defines store cache key store
type IncomingRequestCacheKeyStore struct {
	keys map[string]string
	mu   sync.RWMutex
}

// NewIncomingRequestCacheKeyStore Initiate new cache key store
func NewIncomingRequestCacheKeyStore() *IncomingRequestCacheKeyStore {
	return &IncomingRequestCacheKeyStore{
		keys: make(map[string]string),
	}
}

// Set Incoming request cache key
func (s *IncomingRequestCacheKeyStore) Set(requestID string, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys[requestID] = key
	fmt.Printf("[IncomingRequestCacheKeyStore] SET requestID=%q -> key=%q\n", requestID, key)
	return nil
}

// Pop the request cache key.
func (s *IncomingRequestCacheKeyStore) Pop(requestID string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key, exists := s.keys[requestID]
	fmt.Printf("[IncomingRequestCacheKeyStore] POP requestID=%q -> key=%q, exists=%v\n", requestID, key, exists)
	if !exists {
		return "", false
	}
	delete(s.keys, requestID)
	return key, true
}
