package colors

import (
	"github.com/fatih/color"
	_ "github.com/gookit/color"
)

var (
	FatalPrefix   = color.New(color.BgHiRed, color.FgBlack).SprintFunc()
	SuccessPrefix = color.New(color.BgHiGreen, color.FgBlack).SprintFunc()
	DebugPrefix   = color.New(color.BgHiWhite, color.FgBlack).SprintFunc()
	HiRed         = color.New(color.FgHiRed).SprintFunc()
	HiYellow      = color.New(color.FgHiYellow).SprintFunc()
	HiCyan        = color.New(color.FgHiCyan).SprintFunc()
	HiMagenta     = color.New(color.FgHiMagenta).SprintFunc()
	Bold          = color.New(color.Bold).SprintFunc()
)
