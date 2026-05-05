package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"vital-fitness/backend/internal/config"
)

// DashScope API error definitions
var (
	ErrDashScopeTimeout       = errors.New("识别超时，请稍后重试")
	ErrDashScopeUnavailable   = errors.New("AI 服务暂时不可用，请稍后重试")
	ErrDashScopeNotConfigured = errors.New("AI 识别服务未配置")
)

// DashScopeClient handles communication with the DashScope OpenAI-compatible API.
type DashScopeClient struct {
	apiKey     string
	model      string
	baseURL    string
	httpClient *http.Client
}

// chatCompletionRequest represents the OpenAI-compatible chat completion request body.
type chatCompletionRequest struct {
	Model       string        `json:"model"`
	Messages    []chatMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
}

// chatMessage represents a single message in the chat completion request.
type chatMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

// contentPart represents a part of a multimodal user message.
type contentPart struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *imageURL `json:"image_url,omitempty"`
}

// imageURL represents the image URL object in a content part.
type imageURL struct {
	URL string `json:"url"`
}

// chatCompletionResponse represents the DashScope API response.
type chatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// systemPrompt is the food recognition prompt sent to the model.
const systemPrompt = `你是一个专业的食物营养分析助手。请分析图片中的食物，返回JSON格式的识别结果。

要求：
1. 识别图片中所有可见的食物
2. 为每种食物估算合理的份量（克数）
3. 提供每100g的营养数据（热量、蛋白质、碳水化合物、脂肪）
4. 食物名称使用中文
5. 如果无法识别图片中的食物，返回空数组

请严格按照以下JSON格式返回，不要包含任何其他文字：
{
  "foods": [
    {
      "name": "食物名称",
      "estimated_grams": 150,
      "calories_per_100g": 116,
      "protein_per_100g": 20.5,
      "carbs_per_100g": 0,
      "fat_per_100g": 3.8,
      "confidence": 0.9
    }
  ]
}

confidence 取值 0-1，表示识别置信度。低于 0.6 视为低置信度。`

// NewDashScopeClient creates a new DashScope client with the given configuration.
func NewDashScopeClient(cfg config.DashScope) *DashScopeClient {
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 15
	}

	return &DashScopeClient{
		apiKey:  cfg.APIKey,
		model:   cfg.Model,
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// RecognizeFood sends an image to the DashScope API and returns the raw model response text.
func (c *DashScopeClient) RecognizeFood(ctx context.Context, imageBase64 string, mimeType string) (string, error) {
	if c.apiKey == "" {
		return "", ErrDashScopeNotConfigured
	}

	// Construct the data URL for the image
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, imageBase64)

	// Build the request body
	reqBody := chatCompletionRequest{
		Model: c.model,
		Messages: []chatMessage{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role: "user",
				Content: []contentPart{
					{
						Type:     "image_url",
						ImageURL: &imageURL{URL: dataURL},
					},
					{
						Type: "text",
						Text: "请识别图片中的所有食物并估算营养信息。",
					},
				},
			},
		},
		Temperature: 0.1,
		MaxTokens:   2048,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Build the HTTP request
	url := c.baseURL + "/chat/completions"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Check if the error is a timeout
		if isTimeoutError(err) {
			return "", ErrDashScopeTimeout
		}
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		return "", ErrDashScopeUnavailable
	}

	// Read and parse the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var completionResp chatCompletionResponse
	if err := json.Unmarshal(respBody, &completionResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(completionResp.Choices) == 0 {
		return "", ErrDashScopeUnavailable
	}

	return completionResp.Choices[0].Message.Content, nil
}

// isTimeoutError checks if an error is a timeout error.
func isTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	// Check for net.Error interface with Timeout() method
	type timeoutErr interface {
		Timeout() bool
	}
	var te timeoutErr
	if errors.As(err, &te) {
		return te.Timeout()
	}
	return false
}
