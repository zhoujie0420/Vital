package router

import (
	"net/http"

	"vital-fitness/backend/internal/handler"
	"vital-fitness/backend/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "vital-fitness/backend/docs"
)

// SetupRouter 设置路由
func SetupRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Vital Fitness API is running",
		})
	})

	// 初始化 handler
	authHandler := handler.NewAuthHandler()

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 认证相关接口（无需登录）
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		// 需要认证的接口组
		authorized := v1.Group("/")
		authorized.Use(middleware.JWTAuth())
		{
			// 用户相关接口
			users := authorized.Group("/users")
			{
				users.GET("/profile", authHandler.GetProfile)
				users.PUT("/profile", placeholder("update user profile"))
			}

			// 健身记录接口
			workouts := authorized.Group("/workouts")
			{
				workouts.POST("/", placeholder("create workout"))
				workouts.GET("/", placeholder("get workouts"))
				workouts.GET("/:id", placeholder("get workout"))
				workouts.PUT("/:id", placeholder("update workout"))
				workouts.DELETE("/:id", placeholder("delete workout"))
			}

			// 动作库接口
			exercises := authorized.Group("/exercises")
			{
				exercises.GET("/", placeholder("get exercises"))
				exercises.POST("/", placeholder("create exercise"))
			}

			// 心情记录接口
			moods := authorized.Group("/moods")
			{
				moods.POST("/", placeholder("create mood"))
				moods.GET("/", placeholder("get moods"))
			}

			// 体重记录接口
			weights := authorized.Group("/weights")
			{
				weights.POST("/", placeholder("create weight"))
				weights.GET("/", placeholder("get weights"))
			}

			// 饮食记录接口
			diets := authorized.Group("/diets")
			{
				diets.POST("/", placeholder("create diet"))
				diets.GET("/", placeholder("get diets"))
			}
		}
	}

	// Swagger API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// placeholder 占位处理函数，后续版本替换
func placeholder(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": name + " - 待实现"})
	}
}
