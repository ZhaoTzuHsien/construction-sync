package main

import (
	"fmt"
	"github.com/ZhaoTzuHsien/construction-sync/internal/app/sync"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/colors"
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/log"
)

func main() {
	defer func() {
		log.NoFlag.Println()
		log.NoFlag.Printf(
			"%s is an open source project. Made with %s by %s.",
			colors.HiCyan("Construction sync"),
			colors.HiRed("❤"),
			colors.Bold("趙子賢"))
		fmt.Scanln()
	}()

	sync.Start()
}
