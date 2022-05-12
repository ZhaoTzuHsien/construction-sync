package log

import (
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/colors"
	"log"
	"os"
)

var (
	Default = log.Default()
	Fatal   = log.New(os.Stderr, colors.FatalPrefix("錯誤")+" ", 0)
	Success = log.New(os.Stdout, colors.SuccessPrefix("成功")+" ", 0)
	Debug   = log.New(os.Stdout, colors.DebugPrefix("偵錯")+" ", 0)
	NoFlag  = log.New(os.Stdout, "", 0)
)
