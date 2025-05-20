package router

import (
	"Gocument/controllers"
	"Gocument/middleware"
	"Gocument/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *services.AuthService, documentService *services.DocumentService) *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 认证路由（无需中间件）
	authController := controllers.NewAuthController(authService)
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}

	// 需要认证的路由（应用 JWT 中间件）
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware()) // 应用 JWT 认证中间件
	{
		documentController := controllers.NewDocumentController(documentService)
		documentsGroup := apiGroup.Group("/documents")
		{
			documentsGroup.POST("", documentController.CreateDocument)
			documentsGroup.GET("/:id", documentController.GetDocument)
			documentsGroup.PUT("/:id", documentController.UpdateDocument)
			documentsGroup.DELETE("/:id", documentController.DeleteDocument)
		}
	}

	return r
}
