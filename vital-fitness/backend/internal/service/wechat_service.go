package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"vital-fitness/backend/internal/config"
	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

// WeChatService 微信登录服务
type WeChatService struct {
	userDAO   *dao.UserDAO
	appID     string
	appSecret string
}

// WxSessionResponse 微信 code2session 响应
type WxSessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// NewWeChatService 创建实例
func NewWeChatService(cfg config.WeChat) *WeChatService {
	return &WeChatService{
		userDAO:   dao.NewUserDAO(),
		appID:     cfg.AppID,
		appSecret: cfg.AppSecret,
	}
}

// Code2Session 用 code 换取 openid
func (s *WeChatService) Code2Session(code string) (*WxSessionResponse, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		s.appID, s.appSecret, code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("请求微信接口失败")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取微信响应失败")
	}

	var session WxSessionResponse
	if err := json.Unmarshal(body, &session); err != nil {
		return nil, errors.New("解析微信响应失败")
	}

	if session.ErrCode != 0 {
		return nil, fmt.Errorf("微信登录失败: %s", session.ErrMsg)
	}

	return &session, nil
}

// LoginOrRegister 微信登录（自动注册）
func (s *WeChatService) LoginOrRegister(openID string) (*model.User, error) {
	// 查找已有用户
	user, err := s.userDAO.FindByOpenID(openID)
	if err == nil {
		return user, nil
	}

	// 不存在则自动创建
	user = &model.User{
		OpenID:   openID,
		Username: "wx_" + openID[:8],
		Nickname: "微信用户",
	}

	if err := s.userDAO.Create(user); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}
