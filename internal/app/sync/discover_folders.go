package sync

import (
	"errors"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/log"
	"github.com/spf13/viper"
	"path/filepath"
)

func getSourceDirs() ([]string, error) {
	// Check if source.path exist and retrieve absolute path
	sourcePath := viper.GetString("source.path")
	absSourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		panic(errors.New("無法將 source.path 轉換成絕對路徑"))
	}

	// Build absolute glob
	glob := viper.GetString("source.glob")
	absGlob := filepath.Join(absSourcePath, glob)

	log.Debug.Println("absolute glob:", absGlob)

	return filepath.Glob(absGlob)
}
