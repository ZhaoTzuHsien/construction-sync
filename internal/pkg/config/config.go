package config

import (
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/constant"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfig() {
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
			var errorMsg string
			errorMsg += "找不到設定檔，請確認 config.yaml 是否存在於以下資料夾："
			for _, v := range configs {
				errorMsg += "\n - " + v
			}
			panic(errorMsg)
		}
	}

	// validate config
	notFoundKeys := validate([]string{"source.path", "source.glob", "destination.path"})
	if len(notFoundKeys) > 0 {
		panic("無法在 config.yaml 中找到 " + strings.Join(notFoundKeys, ", "))
	}
}

func validate(keys []string) []string {
	var notFoundKeys []string

	for _, v := range keys {
		if !viper.IsSet(v) {
			notFoundKeys = append(notFoundKeys, v)
		}
	}

	return notFoundKeys
}
