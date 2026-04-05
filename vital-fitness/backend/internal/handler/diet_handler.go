package handler

import (
	"net/http"
	"strconv"

	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type DietHandler struct{ dietService *service.DietService }

func NewDietHandler() *DietHandler { return &DietHandler{dietService: service.NewDietService()} }

func (h *DietHandler) CreateRecord(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req model.CreateDietRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	r, err := h.dietService.CreateRecord(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": r})
}

func (h *DietHandler) GetRecords(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	records, total, err := h.dietService.GetRecords(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": records, "total": total})
}

func (h *DietHandler) GetFoods(c *gin.Context) {
	userID := c.GetUint("user_id")
	keyword := c.Query("keyword")
	category := c.Query("category")
	foods, err := h.dietService.GetFoods(keyword, category, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": foods})
}

func (h *DietHandler) CreateFood(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req model.CreateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误: " + err.Error()})
		return
	}
	f, err := h.dietService.CreateFood(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": f})
}
