package sync

import (
	"errors"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/colors"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/config"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/log"
	"github.com/schollz/progressbar/v3"
	"strings"
	"sync"
)

type Progress struct {
	task   string
	action string
}

func Start() {
	// Error handling
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal.Fatalln(err)
		}
	}()

	configFilePath := config.LoadConfig()
	log.Success.Println("載入", configFilePath)

	sourceDirs, err := getSourceDirs()
	if err != nil {
		panic(errors.New("source.glob 格式錯誤，請參閱以下網址修正\nhttps://en.wikipedia.org/wiki/Glob_(programming)#Syntax"))
	}

	srcDestMap := createSrcDestMap(sourceDirs)
	log.Success.Println("搜尋資料夾，將從以下路徑搜尋檔案：\n - " + strings.Join(sourceDirs, "\n - "))

	// Create channels for discover, hash and copy tasks
	hashChannel := make(chan [2]string, 100)
	copyChannel := make(chan [2]string, 100)
	progressChannel := make(chan Progress, 100*2*3)
	errorChannel := make(chan error)
	done := make(chan struct{})

	// Create task counts
	hashCount, copyCount := 0, 0

	// Create progress bar
	log.NoFlag.Println()
	bar := progressbar.NewOptions(-1,
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("正在複製檔案..."),
		progressbar.OptionSetItsString("檔案"),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSpinnerType(24),
	)

	// Discover files
	discoverFiles(srcDestMap, hashChannel, progressChannel, errorChannel)

	// Handle hashChannel, copyChannel and errorChannel as pipeline
	var wgCheckHash sync.WaitGroup
	var wgCopyFile sync.WaitGroup
	for {
		select {
		case hashPair, ok := <-hashChannel:
			/**
			If hasChannel is closed, prevent this case from executing again and wait for all file check goroutines done.
			Then, close copyChannel.
			*/
			if !ok {
				hashChannel = nil

				go func() {
					wgCheckHash.Wait()
					close(copyChannel)
				}()

				continue
			}

			/**
			Whenever a hashPair is arrived, add one to WaitGroup and spawn a goroutine to check if the source and destination files are the same.
			After that, mark the goroutine done.
			*/
			wgCheckHash.Add(1)
			progressChannel <- Progress{task: "hash", action: "start"}

			go func() {
				defer func() {
					wgCheckHash.Done()
					progressChannel <- Progress{task: "hash", action: "finish"}
				}()

				if same := isSameFile(hashPair[0], hashPair[1]); !same {
					copyChannel <- hashPair
				}
			}()
		case copyPair, ok := <-copyChannel:
			/**
			If copyChannel is closed, prevent this case from executing again and wait for all copy goroutines done.
			Then, close done.
			*/
			if !ok {
				copyChannel = nil

				go func() {
					wgCopyFile.Wait()
					close(done)
				}()

				continue
			}

			/**
			Whenever a copyPair is arrived, add one to WaitGroup and spawn a goroutine to copy file.
			If an error occurs, push that error to errorChannel.
			After that, mark the goroutine done.
			*/
			wgCopyFile.Add(1)
			progressChannel <- Progress{task: "copy", action: "start"}

			go func() {
				defer func() {
					wgCopyFile.Done()
					progressChannel <- Progress{task: "copy", action: "finish"}
				}()
				err := copy(copyPair[0], copyPair[1])
				if err != nil {
					errorChannel <- err
				}
			}()
		// Listen to progress channel and update progress
		case progress := <-progressChannel:
			if progress.action == "finish" {
				switch progress.task {
				case "hash":
					hashCount++
				case "copy":
					bar.Add(1)
					copyCount++
				}
			}
		// Listen to errorChannel and log fatal error content
		case err := <-errorChannel:
			log.Fatal.Fatalln(err)
		// Exit function if both check hash and copy file tasks done
		case <-done:
			bar.Finish()
			log.Success.Printf("複製 %s 個檔案並跳過 %s 個重複的檔案", colors.HiYellow(copyCount), colors.HiMagenta(hashCount-copyCount))
			return
		}
	}
}
