package main

import (
	"fmt"
	"github.com/ZhaoTzuHsien/construction-sync/internal/app/sync"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			_, _ = fmt.Scanln()
		}
	}()

	sync.Start()
}
