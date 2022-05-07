package sync

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return errors.New("無法開啟檔案：" + src)
	}
	defer in.Close()

	// Create parent directory
	err = os.MkdirAll(filepath.Dir(dst), os.ModeDir)
	if err != nil {
		return errors.New("無法建立資料夾：" + filepath.Dir(dst))
	}

	out, err := os.Create(dst)
	if err != nil {
		return errors.New("無法建立檔案：" + dst)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return errors.New(fmt.Sprintf("無法複製 %s\n -> %s", src, dst))
	}
	return out.Close()
}
