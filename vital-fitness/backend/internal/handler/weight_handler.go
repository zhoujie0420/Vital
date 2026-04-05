package handler

import (
	"net/http"
	"strconv"

	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type WeightHandler struct{ weightService *service.WeightService }

func NewWeightHandler() *WeightHandler {
	return &WeightHandler{weightService: service.NewWeightService()}
}

func (h *WeightHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req model.CreateWeightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	r, err := h.weightService.Create(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": r})
}

func (h *WeightHandler) GetList(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	records, total, err := h.weightService.GetList(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": records, "total": total})
}
