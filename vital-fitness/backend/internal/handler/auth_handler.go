package handler

import (
	"net/http"

	"vital-fitness/backend/internal/middleware"
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	userService *service.UserService
}

// NewAuthHandler 创建 AuthHandler 实例
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{userService: service.NewUserService()}
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误: " + err.Error()})
		return
	}

	profile, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册成功", "data": profile})
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}

	user, err := h.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": err.Error()})
		return
	}

	// 生成 token
	token, expiresAt, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": model.LoginResponse{
			Token: token,
			User: model.UserProfile{
				ID:        user.ID,
				Username:  user.Username,
				Email:     user.Email,
				Phone:     user.Phone,
				Avatar:    user.Avatar,
				Nickname:  user.Nickname,
				Gender:    user.Gender,
				Birthday:  user.Birthday,
				Height:    user.Height,
				Weight:    user.Weight,
				CreatedAt: user.CreatedAt,
			},
			ExpiresAt: expiresAt,
		},
	})
}

// Logout 用户登出
func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登出成功"})
}

// GetProfile 获取当前用户资料
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未登录"})
		return
	}

	profile, err := h.userService.GetProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success", "data": profile})
}
