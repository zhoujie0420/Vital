package handler

import (
	"net/http"

	"vital-fitness/backend/internal/config"
	"vital-fitness/backend/internal/middleware"
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
type AuthHandler struct {
	userService   *service.UserService
	wechatService *service.WeChatService
}

// NewAuthHandler 创建 AuthHandler 实例
func NewAuthHandler(wxCfg config.WeChat) *AuthHandler {
	return &AuthHandler{
		userService:   service.NewUserService(),
		wechatService: service.NewWeChatService(wxCfg),
	}
}

// WxLogin 微信小程序登录
func (h *AuthHandler) WxLogin(c *gin.Context) {
	var req model.WxLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请提供code"})
		return
	}

	// 用 code 换 openid
	session, err := h.wechatService.Code2Session(req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// 登录或自动注册
	user, err := h.wechatService.LoginOrRegister(session.OpenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
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
				ID:       user.ID,
				Username: user.Username,
				Avatar:   user.Avatar,
				Nickname: user.Nickname,
				Gender:   user.Gender,
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
