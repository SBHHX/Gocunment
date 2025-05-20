package services

import (
	"Gocument/models"
	"Gocument/repositories"
	"Gocument/utils"
	"errors"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

// NewAuthService 初始化认证服务
func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register 用户注册
func (s *AuthService) Register(username, password, nickname string) error {
	// 检查用户名是否已存在
	if _, err := s.userRepo.FindUserByUsername(username); err == nil {
		return errors.New("用户名已存在")
	}

	// 创建新用户
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Nickname: nickname,
	}

	return s.userRepo.CreateUser(user)
}

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if !utils.ComparePasswords(user.Password, password) {
		return "", errors.New("用户名或密码错误")
	}

	// 生成 JWT
	return utils.GenerateToken(user.ID, user.Username)
}
