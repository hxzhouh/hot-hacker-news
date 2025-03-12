package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config 表示应用配置
type Config struct {
	Database struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"database"`

	RSS struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"rss"`

	App struct {
		ItemLimit int `mapstructure:"item_limit"`
	} `mapstructure:"app"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig(configPath string) (*Config, error) {
	// 确保目录存在
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 设置默认值
	viper.SetDefault("database.path", "data/hackernews.db")
	viper.SetDefault("rss.url", "https://www.daemonology.net/hn-daily/index.rss")
	viper.SetDefault("app.item_limit", 5)

	// 如果配置文件不存在，创建默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := viper.SafeWriteConfig(); err != nil {
			return nil, err
		}
	}

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
