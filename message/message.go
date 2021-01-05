package message

import "github.com/fatih/color"

func Highlight(str string) string {
	highlight := color.New(color.FgCyan).SprintFunc()
	return highlight(str)
}

func Success(str string) string {
	success := color.New(color.FgGreen).SprintFunc()
	return success(str)
}

func Error(str string) string {
	err := color.New(color.FgRed).SprintFunc()
	return err(str)
}

func Warning(str string) string {
	warning := color.New(color.FgYellow).SprintFunc()
	return warning(str)
}
