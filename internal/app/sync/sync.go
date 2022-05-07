package sync

import (
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/config"
	"log"
	"sync"
)

func Start() {
	config.LoadConfig()

	sourceDirs, err := getSourceDirs()
	if err != nil {
		panic("glob 格式錯誤")
	}

	srcDestMap := createSrcDestMap(sourceDirs)

	// Create channels for discover, hash and copy tasks
	hashChannel := make(chan [2]string, 10)
	copyChannel := make(chan [2]string, 10)
	errorChannel := make(chan error)
	done := make(chan struct{})

	// Discover files
	discoverFiles(srcDestMap, hashChannel, errorChannel)

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

			go func() {
				defer wgCheckHash.Done()
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

			go func() {
				defer wgCopyFile.Done()
				err := Copy(copyPair[0], copyPair[1])
				if err != nil {
					errorChannel <- err
				}
			}()
		// Listen to errorChannel and log fatal error content
		case err := <-errorChannel:
			log.Fatalln(err.Error())
		// Exit function if both check hash and copy file tasks done
		case <-done:
			return
		}
	}
}
