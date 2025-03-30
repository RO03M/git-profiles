package utils

import "fmt"

type Color string

const (
	None      Color = "\033[00m"
	Red       Color = "\033[01;31m"
	Green     Color = "\033[01;32m"
	Yellow    Color = "\033[01;33m"
	Purple    Color = "\033[01;35m"
	Cyan      Color = "\033[01;36m"
	White     Color = "\033[01;37m"
	Bold      Color = "\033[1m"
	Underline Color = "\033[4m"
)

func TextColor(text string, color Color) string {
	return fmt.Sprintf("%s%s\033[0m", color, text)
}
