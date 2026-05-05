package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"vital-fitness/backend/internal/config"
)

// Error definitions for food recognition
var (
	ErrInvalidImageFormat = errors.New("仅支持 JPG、PNG 格式的图片")
	ErrImageTooLarge      = errors.New("图片大小不能超过 5MB")
	ErrInvalidBase64      = errors.New("图片数据格式错误")
	ErrParseFailure       = errors.New("识别结果解析失败")
)

// Maximum image size: 5MB
const maxImageSize = 5 * 1024 * 1024

// RecognizedFood 单个识别出的食物
type RecognizedFood struct {
	Name            string  `json:"name"`
	EstimatedGrams  float64 `json:"estimated_grams"`
	CaloriesPer100g float64 `json:"calories_per_100g"`
	ProteinPer100g  float64 `json:"protein_per_100g"`
	CarbsPer100g    float64 `json:"carbs_per_100g"`
	FatPer100g      float64 `json:"fat_per_100g"`
	Confidence      float64 `json:"confidence"`
}

// FoodRecognitionResponse 食物识别响应
type FoodRecognitionResponse struct {
	Foods []RecognizedFood `json:"foods"`
}

// FoodRecognizeRequest 食物识别请求
type FoodRecognizeRequest struct {
	Image string `json:"image" binding:"required"`
}

// FoodRecognitionService orchestrates the food recognition flow.
type FoodRecognitionService struct {
	client *DashScopeClient
}

// NewFoodRecognitionService creates a new FoodRecognitionService with the given config.
func NewFoodRecognitionService(cfg config.DashScope) *FoodRecognitionService {
	return &FoodRecognitionService{
		client: NewDashScopeClient(cfg),
	}
}

// RecognizeFood validates the image and returns structured recognition results.
func (s *FoodRecognitionService) RecognizeFood(ctx context.Context, imageBase64 string, mimeType string) (*FoodRecognitionResponse, error) {
	// Validate image format
	detectedMime, err := validateImageFormat(imageBase64)
	if err != nil {
		return nil, err
	}

	// Use detected MIME type if caller didn't provide one
	if mimeType == "" {
		mimeType = detectedMime
	}

	// Validate image size
	if err := validateImageSize(imageBase64); err != nil {
		return nil, err
	}

	// Call DashScope client
	rawResponse, err := s.client.RecognizeFood(ctx, imageBase64, mimeType)
	if err != nil {
		return nil, err
	}

	// Parse model response
	foods, err := parseModelResponse(rawResponse)
	if err != nil {
		return nil, err
	}

	return &FoodRecognitionResponse{Foods: foods}, nil
}

// parseModelResponse extracts structured food data from the model's text output.
// It handles cases where the model wraps JSON in markdown code blocks.
func parseModelResponse(raw string) ([]RecognizedFood, error) {
	// Try to extract JSON from markdown code block wrapping
	jsonStr := extractJSON(raw)

	// Try to parse as a full response with "foods" wrapper
	var response struct {
		Foods []RecognizedFood `json:"foods"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &response); err == nil {
		return response.Foods, nil
	}

	// Try to parse as a direct array of foods
	var foods []RecognizedFood
	if err := json.Unmarshal([]byte(jsonStr), &foods); err == nil {
		return foods, nil
	}

	return nil, ErrParseFailure
}

// extractJSON attempts to extract JSON content from the raw model output.
// It handles markdown code block wrapping (```json ... ``` or ``` ... ```).
func extractJSON(raw string) string {
	raw = strings.TrimSpace(raw)

	// Check for markdown code block with language tag
	if strings.HasPrefix(raw, "```json") {
		raw = strings.TrimPrefix(raw, "```json")
		if idx := strings.LastIndex(raw, "```"); idx >= 0 {
			raw = raw[:idx]
		}
		return strings.TrimSpace(raw)
	}

	// Check for markdown code block without language tag
	if strings.HasPrefix(raw, "```") {
		raw = strings.TrimPrefix(raw, "```")
		if idx := strings.LastIndex(raw, "```"); idx >= 0 {
			raw = raw[:idx]
		}
		return strings.TrimSpace(raw)
	}

	// Try to find JSON object boundaries in text with extra content around it
	startObj := strings.Index(raw, "{")
	startArr := strings.Index(raw, "[")

	if startObj >= 0 && (startArr < 0 || startObj < startArr) {
		// Find the matching closing brace
		endIdx := strings.LastIndex(raw, "}")
		if endIdx > startObj {
			return raw[startObj : endIdx+1]
		}
	}

	if startArr >= 0 {
		endIdx := strings.LastIndex(raw, "]")
		if endIdx > startArr {
			return raw[startArr : endIdx+1]
		}
	}

	return raw
}

// validateImageFormat decodes the first bytes of the base64 data and checks
// for JPEG magic (FF D8 FF) or PNG magic (89 50 4E 47).
// Returns the detected MIME type or an error.
func validateImageFormat(base64Data string) (string, error) {
	// We only need to decode enough bytes to check the magic number
	// 4 bytes is sufficient for both JPEG and PNG detection
	// Base64 encodes 3 bytes into 4 characters, so 8 characters gives us 6 bytes
	minChars := 8
	if len(base64Data) < minChars {
		return "", ErrInvalidBase64
	}

	// Decode the first few bytes
	decoded, err := base64.StdEncoding.DecodeString(padBase64(base64Data[:minChars]))
	if err != nil {
		// Try with the full string in case of padding issues with the slice
		if len(base64Data) > 16 {
			decoded, err = base64.StdEncoding.DecodeString(padBase64(base64Data[:16]))
		}
		if err != nil {
			return "", ErrInvalidBase64
		}
	}

	if len(decoded) < 4 {
		return "", ErrInvalidBase64
	}

	// Check JPEG magic: FF D8 FF
	if decoded[0] == 0xFF && decoded[1] == 0xD8 && decoded[2] == 0xFF {
		return "image/jpeg", nil
	}

	// Check PNG magic: 89 50 4E 47
	if decoded[0] == 0x89 && decoded[1] == 0x50 && decoded[2] == 0x4E && decoded[3] == 0x47 {
		return "image/png", nil
	}

	return "", ErrInvalidImageFormat
}

// validateImageSize checks that the decoded image size does not exceed 5MB.
func validateImageSize(base64Data string) error {
	// Calculate the decoded size from base64 length
	// Base64 encoding increases size by ~4/3, so decoded size ≈ len * 3/4
	// Account for padding characters
	paddingCount := strings.Count(base64Data[max(0, len(base64Data)-4):], "=")
	decodedSize := (len(base64Data) * 3 / 4) - paddingCount

	if decodedSize > maxImageSize {
		return ErrImageTooLarge
	}

	return nil
}

// padBase64 ensures a base64 string has proper padding.
func padBase64(s string) string {
	switch len(s) % 4 {
	case 2:
		return s + "=="
	case 3:
		return s + "="
	default:
		return s
	}
}
