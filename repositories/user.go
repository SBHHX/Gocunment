package repositories

import (
	"Gocument/models"
	"Gocument/utils"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser 创建用户
func (r *UserRepository) CreateUser(user *models.User) error {
	return utils.DB.Create(user).Error
}

// FindUserByUsername 根据用户名查找用户
func (r *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := utils.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
