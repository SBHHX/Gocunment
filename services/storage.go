package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type StorageService struct {
	RootDir string
}

// DeleteFile 删除文件
func (s *StorageService) DeleteFile(relPath string) error {
	filePath := filepath.Join(s.RootDir, relPath)
	return os.Remove(filePath)
}

func NewStorageService(rootDir string) *StorageService {
	// 确保存储目录存在
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		panic(fmt.Sprintf("创建存储目录失败: %v", err))
	}
	return &StorageService{RootDir: rootDir}
}

// GenerateUniqueFileName 生成唯一文件名
func (s *StorageService) GenerateUniqueFileName(ext string) string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b) + "." + strings.TrimPrefix(ext, ".")
}

// SaveDocumentContent 保存文档内容到文件
func (s *StorageService) SaveDocumentContent(userID uint, content []byte) (string, error) {
	userDir := filepath.Join(s.RootDir, fmt.Sprintf("user_%d", userID))
	if err := os.MkdirAll(userDir, 0755); err != nil {
		return "", err
	}

	fileName := s.GenerateUniqueFileName("txt")
	filePath := filepath.Join(userDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.WriteString(file, string(content))
	if err != nil {
		return "", err
	}

	// 返回相对于根目录的路径
	relPath, _ := filepath.Rel(s.RootDir, filePath)
	return relPath, nil
}

// ReadDocumentContent 读取文档内容
func (s *StorageService) ReadDocumentContent(relPath string) ([]byte, error) {
	filePath := filepath.Join(s.RootDir, relPath)
	return os.ReadFile(filePath)
}
