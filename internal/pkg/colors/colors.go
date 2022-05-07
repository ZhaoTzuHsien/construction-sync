package colors

import "github.com/fatih/color"

var (
	FatalPrefix   = color.New(color.BgHiRed, color.FgBlack).SprintFunc()
	SuccessPrefix = color.New(color.BgHiGreen, color.FgBlack).SprintFunc()
)
