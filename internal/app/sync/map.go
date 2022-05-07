package sync

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
	"strings"
)

func createSrcDestMap(sourceDirs []string) map[string]string {
	// Get absolute destination path
	destPath := viper.GetString("destination.path")
	absDestPath, err := filepath.Abs(destPath)
	if err != nil {
		panic(errors.New("無法將 destination.path 轉換成絕對路徑"))
	}

	// Create source destination map
	var srcDestMap = make(map[string]string)
	const format = "單元 (%d)"
	for _, v := range sourceDirs {
		srcDirName := filepath.Base(v)
		no := retrieveNo(srcDirName)
		// 單元 ...
		destParentDirName := fmt.Sprintf(format, no)

		srcDestMap[v] = filepath.Join(absDestPath, destParentDirName, srcDirName)
	}

	return srcDestMap
}

func retrieveNo(dirName string) int {
	fullNoStr := strings.Split(dirName, " ")[1]
	noStr := strings.Split(fullNoStr, ".")[1]
	if no, err := strconv.Atoi(noStr); err == nil {
		return no
	} else {
		panic(errors.New("無法將編號從 string 轉換成 int"))
	}
}
