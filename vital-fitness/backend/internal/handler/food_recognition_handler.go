package handler

import (
	"errors"
	"net/http"

	"vital-fitness/backend/internal/config"
	"vital-fitness/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// FoodRecognitionHandler handles food recognition HTTP requests.
type FoodRecognitionHandler struct {
	service *service.FoodRecognitionService
}

// NewFoodRecognitionHandler creates a new FoodRecognitionHandler with the given DashScope config.
func NewFoodRecognitionHandler(cfg config.DashScope) *FoodRecognitionHandler {
	return &FoodRecognitionHandler{
		service: service.NewFoodRecognitionService(cfg),
	}
}

// RecognizeFood handles POST /api/v1/foods/recognize
func (h *FoodRecognitionHandler) RecognizeFood(c *gin.Context) {
	var req service.FoodRecognizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请上传食物图片"})
		return
	}

	// Call the service to validate and recognize
	result, err := h.service.RecognizeFood(c.Request.Context(), req.Image, "")
	if err != nil {
		h.handleError(c, err)
		return
	}

	// Check if no food was recognized (empty array)
	if len(result.Foods) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "未能识别出食物，请尝试重新拍照", "data": result})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "识别成功", "data": result})
}

// handleError maps service errors to appropriate HTTP responses.
func (h *FoodRecognitionHandler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidImageFormat):
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "仅支持 JPG、PNG 格式的图片"})
	case errors.Is(err, service.ErrImageTooLarge):
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "图片大小不能超过 5MB"})
	case errors.Is(err, service.ErrInvalidBase64):
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "图片数据格式错误"})
	case errors.Is(err, service.ErrDashScopeTimeout):
		c.JSON(http.StatusGatewayTimeout, gin.H{"code": 504, "message": "识别超时，请稍后重试"})
	case errors.Is(err, service.ErrDashScopeUnavailable):
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "message": "AI 服务暂时不可用，请稍后重试"})
	case errors.Is(err, service.ErrDashScopeNotConfigured):
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "AI 识别服务未配置"})
	case errors.Is(err, service.ErrParseFailure):
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "识别结果解析失败"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "服务器内部错误"})
	}
}
