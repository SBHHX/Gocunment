package services

import (
	"Gocument/models"
	"Gocument/repositories"
	"errors"
)

type DocumentService struct {
	repo           *repositories.DocumentRepository
	storageService *StorageService
}

// NewDocumentService 初始化文档服务
func NewDocumentService(repo *repositories.DocumentRepository, storageService *StorageService) *DocumentService {
	return &DocumentService{repo: repo, storageService: storageService}
}

// CreateDocument 创建文档
func (s *DocumentService) CreateDocument(userID uint, title string, content []byte, isPublic bool) error {
	// 保存文档内容到存储
	contentPath, err := s.storageService.SaveDocumentContent(userID, content)
	if err != nil {
		return err
	}

	// 创建文档记录
	document := &models.Document{
		UserID:      userID,
		Title:       title,
		ContentPath: contentPath,
		IsPublic:    isPublic,
	}

	return s.repo.CreateDocument(document)
}

// GetDocument 获取文档
func (s *DocumentService) GetDocument(id uint, userID uint) (*models.Document, error) {
	document, err := s.repo.FindDocumentByID(id)
	if err != nil {
		return nil, err
	}

	// 检查权限
	if !document.IsPublic && document.UserID != userID {
		return nil, errors.New("无权访问该文档")
	}

	return document, nil
}

// services/document.go
// UpdateDocument 更新文档
func (s *DocumentService) UpdateDocument(id uint, userID uint, title string, content []byte, isPublic bool) error {
	// 正确接收两个返回值，此时 document 类型为 *models.Document
	document, err := s.repo.FindDocumentByID(id)
	if err != nil {
		return err
	}

	// 检查权限
	if document.UserID != userID {
		return errors.New("无权更新该文档")
	}

	// 更新内容（假设逻辑）
	if len(content) > 0 {
		newContentPath, err := s.storageService.SaveDocumentContent(userID, content)
		if err != nil {
			return err
		}
		document.ContentPath = newContentPath
	}
	if title != "" {
		document.Title = title
	}
	document.IsPublic = isPublic

	return s.repo.UpdateDocument(document) // 直接传递 document，此时类型为 *models.Document，与方法参数匹配
}

// DeleteDocument 删除文档
func (s *DocumentService) DeleteDocument(id uint, userID uint) error {
	// 接收两个返回值
	document, err := s.repo.FindDocumentByID(id)
	if err != nil {
		return err
	}

	// 检查权限
	if document.UserID != userID {
		return errors.New("无权删除该文档")
	}

	// 删除存储文件（假设逻辑）
	if err := s.storageService.DeleteFile(document.ContentPath); err != nil {
		return err
	}

	// 删除数据库记录
	return s.repo.DeleteDocument(document)
}
