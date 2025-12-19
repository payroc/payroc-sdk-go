package core

import (
	"sync"
	"time"
)

const (
	expirationBufferMinutes          = 2
	defaultInferredAuthExpirySeconds = 3600
)

// InferredAuthProvider manages authentication tokens with caching and automatic refresh.
type InferredAuthProvider struct {
	mu          sync.Mutex
	accessToken string
	expiresAt   time.Time
	fetchMu     sync.Mutex
}

func NewInferredAuthProvider() *InferredAuthProvider {
	return &InferredAuthProvider{}
}

func (i *InferredAuthProvider) SetToken(accessToken string, expiresIn int) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.accessToken = accessToken
	if expiresIn > 0 {
		bufferSeconds := expirationBufferMinutes * 60
		effectiveExpiresIn := max(expiresIn-bufferSeconds, 0)
		i.expiresAt = time.Now().Add(time.Duration(effectiveExpiresIn) * time.Second)
	} else {
		effectiveExpiresIn := max(defaultInferredAuthExpirySeconds-(expirationBufferMinutes*60), 0)
		i.expiresAt = time.Now().Add(time.Duration(effectiveExpiresIn) * time.Second)
	}
}

func (i *InferredAuthProvider) GetToken() string {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.accessToken == "" {
		return ""
	}
	if !i.expiresAt.IsZero() && time.Now().After(i.expiresAt) {
		return ""
	}
	return i.accessToken
}

func (i *InferredAuthProvider) GetOrFetch(fetchFunc func() (string, int, error)) (string, error) {
	if token := i.GetToken(); token != "" {
		return token, nil
	}

	i.fetchMu.Lock()
	defer i.fetchMu.Unlock()

	if token := i.GetToken(); token != "" {
		return token, nil
	}

	accessToken, expiresIn, err := fetchFunc()
	if err != nil {
		return "", err
	}

	i.SetToken(accessToken, expiresIn)
	return accessToken, nil
}

func (i *InferredAuthProvider) NeedsRefresh() bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.accessToken == "" {
		return true
	}
	if !i.expiresAt.IsZero() && time.Now().After(i.expiresAt) {
		return true
	}
	return false
}

func (i *InferredAuthProvider) Reset() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.accessToken = ""
	i.expiresAt = time.Time{}
}
