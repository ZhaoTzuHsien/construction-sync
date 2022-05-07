package colors

import "github.com/fatih/color"

var (
	FatalPrefix   = color.New(color.BgHiRed, color.FgBlack).SprintFunc()
	SuccessPrefix = color.New(color.BgHiGreen, color.FgBlack).SprintFunc()
	HiYellow      = color.New(color.FgHiYellow).SprintFunc()
	HiCyan        = color.New(color.FgHiCyan).SprintFunc()
	HiMagenta     = color.New(color.FgHiMagenta).SprintFunc()
)
