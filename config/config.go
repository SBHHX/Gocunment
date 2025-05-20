package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	Storage struct {
		RootDir string `yaml:"root_dir"`
		TempDir string `yaml:"temp_dir"`
	} `yaml:"storage"`
	JWT struct {
		Secret    string `json:"secret" yaml:"secret"`
		ExpiresIn int    `json:"expires_in" yaml:"expires_in"`
	} `yaml:"jwt"`
}

// LoadConfig 从配置文件和环境变量加载配置
func LoadConfig() (*Config, error) {
	// 默认配置路径
	configPath := "config/config.yaml"

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 支持通过环境变量覆盖配置
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		cfg.Server.Port = envPort
	}

	if envMode := os.Getenv("SERVER_MODE"); envMode != "" {
		cfg.Server.Mode = envMode
	}

	if envDSN := os.Getenv("DB_DSN"); envDSN != "" {
		cfg.Database.DSN = envDSN
	}

	if envRootDir := os.Getenv("STORAGE_ROOT_DIR"); envRootDir != "" {
		cfg.Storage.RootDir = envRootDir
	}

	if envJWTSecret := os.Getenv("JWT_SECRET"); envJWTSecret != "" {
		cfg.JWT.Secret = envJWTSecret
	}

	return &cfg, nil
}
