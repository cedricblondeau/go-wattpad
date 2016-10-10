package wattpad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	apiKey = "1234567890"
)

func TestNewClient(t *testing.T) {
	c := NewClient(apiKey)
	assert.Equal(t, apiKey, c.Key)
}

func TestNewRequest(t *testing.T) {
	c := NewClient(apiKey)
	endpoint := "stories/123"
	expectedURL := apiURL + "/" + endpoint
	method := "GET"
	req := c.buildRequest(method, endpoint)
	assert.Equal(t, method, req.Method)
	assert.Equal(t, userAgent, req.UserAgent())
	assert.Equal(t, expectedURL, req.URL.String())
}

func TestNewRequestWithPost(t *testing.T) {
	c := NewClient(apiKey)
	method := "POST"
	req := c.buildRequest(method, "stories")
	assert.Equal(t, method, req.Method)
}

func TestNewRequestWithParams(t *testing.T) {
	c := NewClient(apiKey)
	req := c.buildRequest("GET", "stories", param{name: "hello", value: "world"})
	expectedURL := apiURL + "/stories?hello=world"
	assert.Equal(t, expectedURL, req.URL.String())
}
