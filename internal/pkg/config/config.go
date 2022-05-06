package config

import (
	"fmt"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/constant"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	// Define config paths
	var configs []string
	if workDir, err := os.Getwd(); err == nil {
		configs = append(configs, workDir)
	}
	if localConfigDir, err := filepath.Abs("./configs"); err == nil {
		configs = append(configs, localConfigDir)
	}
	if userConfigDir, err := os.UserConfigDir(); err == nil {
		configs = append(configs, filepath.Join(userConfigDir, constant.APP_NAME))
	}

	// Configure viper with config name, type and location
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	for _, v := range configs {
		viper.AddConfigPath(v)
	}

	// Set Default config
	viper.SetDefault("source.glob", "???年??月/*/*")

	// Read config from local filesystem.
	// If config is not found, ask user to add config.yaml to valid config path
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到設定檔，請確認 config.yaml 是否存在於以下資料夾：")
			for _, v := range configs {
				fmt.Printf(" - %s\n", v)
			}
		}
	}
}
