package main

import (
	"Gocument/config"
	"Gocument/repositories"
	"Gocument/router"
	"Gocument/services"
	"Gocument/utils"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	if err := utils.InitDB(cfg.Database.DSN); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 初始化存储服务
	storageService := services.NewStorageService(cfg.Storage.RootDir)

	// 初始化仓库
	userRepo := repositories.NewUserRepository()         // 返回 *repositories.UserRepository
	documentRepo := repositories.NewDocumentRepository() // 返回 *repositories.DocumentRepository

	// 初始化服务
	authService := services.NewAuthService(userRepo)                             // 传递 *repositories.UserRepository
	documentService := services.NewDocumentService(documentRepo, storageService) // 传递 *repositories.DocumentRepository

	// 设置路由（注意：中间件已在 router 中应用）
	r := router.SetupRouter(authService, documentService)

	// 启动服务器
	port := cfg.Server.Port
	log.Printf("服务器启动在端口: %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
