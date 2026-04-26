package handler

import (
	"net/http"
	"time"

	"vital-fitness/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct{}

func NewStatsHandler() *StatsHandler { return &StatsHandler{} }

// GetDashboard 首页仪表盘数据
func (h *StatsHandler) GetDashboard(c *gin.Context) {
	userID := c.GetUint("user_id")
	db := utils.GetDB()
	today := time.Now().Format("2006-01-02")

	var workoutCount int64
	db.Table("workouts").Where("user_id = ? AND deleted_at IS NULL", userID).Count(&workoutCount)

	var todayWorkouts int64
	db.Table("workouts").Where("user_id = ? AND DATE(workout_date) = ? AND deleted_at IS NULL", userID, today).Count(&todayWorkouts)

	var latestWeight float64
	db.Table("weight_records").Where("user_id = ? AND deleted_at IS NULL", userID).
		Order("record_date DESC").Limit(1).Pluck("weight", &latestWeight)

	var todayCalories float64
	db.Table("diet_records").Where("user_id = ? AND DATE(record_date) = ? AND deleted_at IS NULL", userID, today).
		Pluck("COALESCE(SUM(total_calories), 0)", &todayCalories)

	var todayDietCount int64
	db.Table("diet_records").Where("user_id = ? AND DATE(record_date) = ? AND deleted_at IS NULL", userID, today).Count(&todayDietCount)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"workout_count":    workoutCount,
			"today_workouts":   todayWorkouts,
			"latest_weight":    latestWeight,
			"today_calories":   todayCalories,
			"today_diet_count": todayDietCount,
		},
	})
}

// GetStats 统计页数据
func (h *StatsHandler) GetStats(c *gin.Context) {
	userID := c.GetUint("user_id")
	db := utils.GetDB()
	period := c.DefaultQuery("period", "week")

	var interval string
	switch period {
	case "month":
		interval = "30 DAY"
	case "year":
		interval = "365 DAY"
	default:
		interval = "7 DAY"
	}

	// 训练次数
	type DayStat struct {
		Day   string `json:"day"`
		Count int    `json:"count"`
	}
	var workoutStats []DayStat
	db.Raw(`SELECT DATE(workout_date) as day, COUNT(*) as count 
		FROM workouts WHERE user_id = ? AND deleted_at IS NULL 
		AND workout_date >= DATE_SUB(CURDATE(), INTERVAL `+interval+`)
		GROUP BY DATE(workout_date) ORDER BY day`, userID).Scan(&workoutStats)

	// 体重趋势
	type WeightPoint struct {
		Day    string  `json:"day"`
		Weight float64 `json:"weight"`
	}
	var weightTrend []WeightPoint
	db.Raw(`SELECT DATE(record_date) as day, weight 
		FROM weight_records WHERE user_id = ? AND deleted_at IS NULL 
		AND record_date >= DATE_SUB(CURDATE(), INTERVAL `+interval+`)
		ORDER BY record_date ASC`, userID).Scan(&weightTrend)

	// 饮食热量
	type CalStat struct {
		Day      string  `json:"day"`
		Calories float64 `json:"calories"`
	}
	var dietStats []CalStat
	db.Raw(`SELECT DATE(record_date) as day, SUM(total_calories) as calories 
		FROM diet_records WHERE user_id = ? AND deleted_at IS NULL 
		AND record_date >= DATE_SUB(CURDATE(), INTERVAL `+interval+`)
		GROUP BY DATE(record_date) ORDER BY day`, userID).Scan(&dietStats)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"workout_stats": workoutStats,
			"weight_trend":  weightTrend,
			"diet_stats":    dietStats,
		},
	})
}
