package log

import (
	"github.com/ZhaoTzuHsien/construction-sync/internal/pkg/colors"
	"log"
	"os"
)

var (
	Fatal   = log.New(os.Stderr, colors.FatalPrefix("錯誤")+" ", 0)
	Success = log.New(os.Stdout, colors.SuccessPrefix("成功")+" ", 0)
)
