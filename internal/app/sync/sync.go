package sync

import (
	"fmt"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/config"
	"github.com/spf13/viper"
	"path/filepath"
)

func Start() {
	config.LoadConfig()

	sourceDirs, err := getSourceDirs()
	if err != nil {
		panic("glob 格是錯誤")
	}
	for _, v := range sourceDirs {
		fmt.Println(v)
	}
}

func getSourceDirs() ([]string, error) {
	// Check if source.path exist and retrieve absolute path
	sourcePath := viper.GetString("source.path")
	absSourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		panic("Cannot get absolute path of source.path.")
	}

	// Build absolute glob
	glob := viper.GetString("source.glob")
	absGlob := filepath.Join(absSourcePath, glob)

	return filepath.Glob(absGlob)
}
