package sync

import (
	"fmt"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/utils"
	"io"
	"os"
	"path/filepath"
)

func copyFiles(srcDestMap map[string]string) {
	hashChannel := make(chan [2]string, 1000)
	copyChannel := make(chan [2]string, 1000)

	var cs []chan bool
	for k, v := range srcDestMap {
		c := pushToHash(k, v, hashChannel)
		cs = append(cs, c)
	}

	// Wait for all pushToHash ended and close hash channel
	go func() {
		for _, c := range cs {
			for range c {
			}
		}
		close(hashChannel)
	}()

	go func() {
		var cs []chan byte
		for pair := range hashChannel {
			c := checkHashAndPush(pair, copyChannel)
			cs = append(cs, c)
		}

		// Wait for all checkHashAndPush ended and close copy channel
		go func() {
			for _, c := range cs {
				for range c {
				}
			}
			close(copyChannel)
		}()
	}()

	for pair := range copyChannel {
		pair := pair
		go func() {
			err := Copy(pair[0], pair[1])
			if err != nil {
				fmt.Println(err)
				fmt.Printf("無法將 %s 複製到 %s", pair[0], pair[1])
			}
		}()
	}
}

// Generate actual source file and destination file name pairs to hash
func pushToHash(sourceDir string, destinationDir string, hashChannel chan [2]string) chan bool {
	c := make(chan bool)

	go func() {
		file, err := os.Open(sourceDir)
		if err != nil {
			panic("無法開啟資料夾： " + sourceDir)
		}
		defer file.Close()

		entries, err := file.ReadDir(0)
		if err != nil {
			panic("無法讀取資料夾內容： " + sourceDir)
		}

		for _, v := range entries {
			hashChannel <- [2]string{filepath.Join(sourceDir, v.Name()), filepath.Join(destinationDir, v.Name())}
		}

		close(c)
	}()

	return c
}

func checkHashAndPush(srcDestPair [2]string, channel chan [2]string) chan byte {
	c := make(chan byte)

	go func() {
		srcSHA, srcErr := utils.FileSha256(srcDestPair[0])
		destSHA, destErr := utils.FileSha256(srcDestPair[1])
		if srcErr == nil && (destErr != nil || srcSHA != destSHA) {
			channel <- srcDestPair
		}
		close(c)
	}()

	return c
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Create parent directory
	err = os.MkdirAll(filepath.Dir(dst), os.ModeDir)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
