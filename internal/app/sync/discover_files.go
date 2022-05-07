package sync

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

func discoverFiles(srcDestMap map[string]string, channel chan<- [2]string, progressChannel chan<- Progress, errorChannel chan<- error) {
	var wg sync.WaitGroup
	wg.Add(len(srcDestMap))
	for k, v := range srcDestMap {
		progressChannel <- Progress{task: "discover", action: "start"}

		go func(source, dest string) {
			defer func() {
				wg.Done()
				progressChannel <- Progress{task: "discover", action: "finish"}
			}()

			pairs, err := listDir(source, dest)
			if err == nil {
				for _, pair := range pairs {
					channel <- pair
				}
			} else {
				errorChannel <- err
			}
		}(k, v)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()
}

func listDir(sourceDir, destinationDir string) ([][2]string, error) {
	file, err := os.Open(sourceDir)
	if err != nil {
		return nil, errors.New("無法開啟資料夾： " + sourceDir)
	}
	defer file.Close()

	entries, err := file.ReadDir(0)
	if err != nil {
		return nil, errors.New("無法讀取資料夾內容： " + sourceDir)
	}

	var pairs [][2]string
	for _, v := range entries {
		if !v.IsDir() {
			pairs = append(pairs, [2]string{filepath.Join(sourceDir, v.Name()), filepath.Join(destinationDir, v.Name())})
		}
	}

	return pairs, nil
}
