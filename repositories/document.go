package repositories

import (
	"Gocument/models"
	"Gocument/utils"
)

type DocumentRepository struct{}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{}
}

// CreateDocument 创建文档
func (r *DocumentRepository) CreateDocument(document *models.Document) error {
	return utils.DB.Create(document).Error
}

// FindDocumentByID 根据ID查找文档（预加载关联用户）
func (r *DocumentRepository) FindDocumentByID(id uint) (*models.Document, error) {
	var document models.Document
	if err := utils.DB.Preload("User").First(&document, id).Error; err != nil {
		return nil, err
	}
	return &document, nil
}

// UpdateDocument 更新文档
func (r *DocumentRepository) UpdateDocument(document *models.Document) error {
	return utils.DB.Save(document).Error
}

// DeleteDocument 删除文档
func (r *DocumentRepository) DeleteDocument(document *models.Document) error {
	return utils.DB.Delete(document).Error
}
