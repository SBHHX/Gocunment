package controllers

import (
	"net/http"
	"strconv"

	"Gocument/services"
	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService *services.DocumentService
}

func NewDocumentController(documentService *services.DocumentService) *DocumentController {
	return &DocumentController{documentService: documentService}
}

// CreateDocument 创建文档
func (c *DocumentController) CreateDocument(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var request struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		IsPublic bool   `json:"is_public"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.documentService.CreateDocument(
		userID.(uint),
		request.Title,
		[]byte(request.Content),
		request.IsPublic,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建文档失败", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "文档创建成功"})
}

// GetDocument 获取文档
func (c *DocumentController) GetDocument(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文档ID"})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		userID = uint(0) // 未登录用户
	}

	document, err := c.documentService.GetDocument(uint(id), userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文档不存在或无权访问"})
		return
	}

	ctx.JSON(http.StatusOK, document)
}

// UpdateDocument 更新文档
func (c *DocumentController) UpdateDocument(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文档ID"})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未认证"})
		return
	}

	var request struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		IsPublic bool   `json:"is_public"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var content []byte
	if request.Content != "" {
		content = []byte(request.Content)
	}

	if err := c.documentService.UpdateDocument(
		uint(id),
		userID.(uint),
		request.Title,
		content,
		request.IsPublic,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新文档失败", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "文档更新成功"})
}

// DeleteDocument 删除文档
func (c *DocumentController) DeleteDocument(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文档ID"})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未认证"})
		return
	}

	if err := c.documentService.DeleteDocument(uint(id), userID.(uint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除文档失败", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "文档删除成功"})
}
